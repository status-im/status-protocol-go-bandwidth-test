package main

import (
	"context"
	"crypto/ecdsa"
	"database/sql"
	"encoding/hex"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	gonode "github.com/status-im/status-go/node"
	params "github.com/status-im/status-go/params"
	status "github.com/status-im/status-protocol-go"
	gethbridge "github.com/status-im/status-protocol-go/bridge/geth"
	v1 "github.com/status-im/status-protocol-go/v1"
)

type Bstatus struct {
	// message fetching loop controls
	fetchInterval time.Duration
	fetchTimeout  time.Duration
	fetchDone     chan bool

	// Whisper node settings
	whisperDataDir string

	privateKey *ecdsa.PrivateKey  // secret for Status chat identity
	nodeConfig *params.NodeConfig // configuration for Whisper node
	statusNode *gonode.StatusNode // Ethereum Whisper node to run in background
	messenger  *status.Messenger  // Status messaging layer instance

	sourceDir      string
	destinationDir string
}

func (b *Bstatus) Connect(id, addr string, datasync, discovery bool) error {
	key, err := crypto.GenerateKey()
	if err != nil {
		return err
	}
	b.privateKey = key

	b.nodeConfig = b.generateConfig(id, addr)
	b.statusNode = gonode.New()

	accsMgr, _ := b.statusNode.AccountManager()

	if err := b.statusNode.Start(b.nodeConfig, accsMgr); err != nil {
		return err
	}

	shhService, err := b.statusNode.WhisperService()
	if err != nil {
		return err
	}

	// Using an in-memory SQLite DB since we have nothing worth preserving
	db, _ := sql.Open("sqlite3", "file:mem?mode=memory&cache=shared")
	options := []status.Option{
		status.WithDatabase(db),
		//status.WithSendV1Messages(),
	}

	if datasync {
		options = append(options, status.WithDatasync())
	}

	if discovery {
		options = append(options, status.WithGenericDiscoveryTopicSupport())
	}

	messenger, err := status.NewMessenger(
		b.privateKey,
		gethbridge.NewGethWhisperWrapper(shhService),
		"test-1",
		options...,
	)
	if err != nil {
		return err
	}
	b.messenger = messenger

	go b.fetchMessagesLoop()

	return crypto.SaveECDSA(b.sourceDir+"key.txt", key)
}

func (b *Bstatus) Disconnect() error {
	b.stopMessagesLoops()
	if err := b.messenger.Shutdown(); err != nil {
		return err
	}
	if err := b.statusNode.Stop(); err != nil {
		return err
	}
	return nil
}

func (b *Bstatus) Send(chatID string, payload []byte) (string, error) {
	if !b.Connected() {
		return "", fmt.Errorf("Not connected")
	}

	// Use a timeout for sending messages
	ctx, cancel := context.WithTimeout(context.Background(), b.fetchTimeout)
	defer cancel()

	msgHash, err := b.messenger.Send(ctx, chatID, payload)
	if err != nil {
		return "", err
	}
	// TODO handle the delivery event?
	return fmt.Sprintf("%#x", msgHash), nil
}

func (b *Bstatus) Connected() bool {
	return b.statusNode.IsRunning()
}

func (b *Bstatus) generateConfig(id string, addr string) *params.NodeConfig {
	options := []params.Option{
		params.WithFleet(params.FleetStaging),
		b.withListenAddr(addr),
	}

	var configFiles []string
	config, err := params.NewNodeConfigWithDefaultsAndFiles(
		"/tmp/"+id+"geth",
		params.MainNetworkID,
		options,
		configFiles,
	)
	if err != nil {
	}
	return config
}

func (b *Bstatus) stopMessagesLoops() {
	close(b.fetchDone)
}

func (b *Bstatus) fetchMessagesLoop() {
	privateRead, err := os.Create(b.sourceDir + "private-read.txt")
	if err != nil {
		fmt.Printf("Error creating private file: %+v", err)
		return
	}

	t := time.NewTicker(b.fetchInterval)
	defer t.Stop()
	for {
		select {
		case <-t.C:
			messages, err := b.retrieveLatestMessages()
			if err != nil {
				continue
			}
			for _, msg := range messages {
				privateRead.WriteString("0x" + hex.EncodeToString(msg.ID) + "\n")
			}
		case <-b.fetchDone:
			return
		}
	}
}

func (b *Bstatus) retrieveLatestMessages() ([]*v1.StatusMessage, error) {
	var msgs []*v1.StatusMessage
	rawMessages, err := b.messenger.RetrieveRawAll()
	if err != nil {
		return nil, err
	}
	for _, msg := range rawMessages {
		msgs = append(msgs, msg...)
	}
	return msgs, nil
}

// skipMessage skips messages that need to be skipped
func (b *Bstatus) skipMessage(msg *v1.Message) bool {
	// skip messages from ourselves
	if isPubKeyEqual(msg.SigPubKey, &b.privateKey.PublicKey) {
		return true
	}

	// skip empty messages
	if msg.Text == "" {
		return true
	}

	return false
}

func publicKeyToHex(pubkey *ecdsa.PublicKey) string {
	return "0x" + hex.EncodeToString(crypto.FromECDSAPub(pubkey))
}

// isPubKeyEqual checks that two public keys are equal
func isPubKeyEqual(a, b *ecdsa.PublicKey) bool {
	// the curve is always the same, just compare the points
	return a.X.Cmp(b.X) == 0 && a.Y.Cmp(b.Y) == 0
}

func (b *Bstatus) JoinChannel(name string) error {
	chat := status.CreatePublicChat(name)
	b.messenger.Join(chat)
	b.messenger.SaveChat(chat)
	return nil
}

func (b *Bstatus) CreateOneToOne(name string, publicKey *ecdsa.PublicKey) error {
	chat := status.CreateOneToOneChat(name, publicKey)
	b.messenger.SaveChat(chat)
	return nil
}

type Destination struct {
	id     string
	key    *ecdsa.PrivateKey
	chatID string
}

func main() {
	src := flag.String("src", "application-1", "this application id")
	dst := flag.String("dst", "application-2", "this application id")
	numberOfMessages := flag.Int("messages", 0, "the number of messages to send")
	numberOfSeconds := flag.Int("seconds", 0, "the number of senconds to run the simulation")
	publicChatID := flag.String("public-chat-id", "test200", "The public chat id to publish messages")
	port := flag.Int("port", 30303, "The port to run geth on")
	datasync := flag.Bool("datasync", false, "Enable datasync")
	discoveryTopic := flag.Bool("discovery", false, "Enabled discovery")

	flag.Parse()

	now := time.Now()
	until := now.Add(time.Duration(*numberOfSeconds) * time.Second)

	addr := fmt.Sprintf("[::]:%d", *port)

	fmt.Printf("Src: %s, Dst: %s, NumberOfMessages: %d, NumberOfSeconds: %d, datasync: %s, discovery: %sPort: %d\n", *src, *dst, *numberOfMessages, *numberOfSeconds, *datasync, *discoveryTopic, *port)

	sourceDir := "/tmp/" + *src + "/"

	os.MkdirAll(sourceDir, os.ModePerm)

	node := &Bstatus{
		sourceDir:     sourceDir,
		fetchInterval: 100 * time.Millisecond,
		fetchTimeout:  1 * time.Second,
	}
	if err := node.Connect(*src, addr, *datasync, *discoveryTopic); err != nil {
		fmt.Printf("Error connecting: %+v", err)
		return
	}

	// Wait a bit, just to make sure all is ready
	time.Sleep(2 * time.Second)

	if *publicChatID != "" {
		if err := node.JoinChannel(*publicChatID); err != nil {
			fmt.Printf("Error connecting: %+v", err)
			return
		}
	}

	publicWrite, err := os.Create(sourceDir + "public-write.txt")
	if err != nil {
		fmt.Printf("Error creating public file: %+v", err)
		return
	}

	if _, err = node.messenger.LoadFilters(nil); err != nil {
		fmt.Printf("Error loading filters: %+v", err)
		return

	}

	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
	sentMessages := 0
	for {
		if *publicChatID != "" {
			id1, err := node.Send(*publicChatID, []byte(fmt.Sprintf("This is one line text message: %d", sentMessages)))

			if err != nil {
				fmt.Printf("Error connecting: %+v", err)
				return
			}

			publicWrite.WriteString(id1 + "\n")
		}
		time.Sleep(300 * time.Millisecond)

		if *numberOfMessages != 0 {
			sentMessages += 1
			if sentMessages == *numberOfMessages {
				break
			}
		}

		if *numberOfSeconds != 0 {
			if until.Before(time.Now()) {
				break
			}
		}

	}

}

func (b *Bstatus) withListenAddr(addr string) params.Option {
	return func(c *params.NodeConfig) error {
		c.ListenAddr = addr
		return nil
	}
}
