package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"

	"github.com/nlopes/slack"
)

func main() {
	slackToken := os.Getenv("SLACK_TOKEN")
	if slackToken == "" {
		log.Fatal("SLACK_TOKEN not set")
	}

	master := os.Getenv("OBEY_USER")
	if master == "" {
		log.Fatal("OBEY_USER not set")
	}

	kubeScript := os.Getenv("KUBE_SCRIPT")
	if kubeScript == "" {
		log.Println("KUBE_SCRIPT not set")
	}

	mineScript := os.Getenv("MINE_SCRIPT")
	if mineScript == "" {
		log.Println("MINE_SCRIPT not set")
	}

	mineIpScript := os.Getenv("MINEIP_SCRIPT")
	if mineIpScript == "" {
		log.Println("MINEIP_SCRIPT not set")
	}

	api := slack.New(slackToken)
	logger := log.New(os.Stdout, "slack-bot: ", log.Lshortfile|log.LstdFlags)
	slack.SetLogger(logger)
	api.SetDebug(false)

	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		go func() {
			var wg sync.WaitGroup

			switch ev := msg.Data.(type) {

			case *slack.MessageEvent:
				var username string
				userinfo, err := api.GetUserInfo(ev.User)
				if err != nil {
					username = ev.Username
				} else {
					username = userinfo.Name
				}
				fmt.Printf("[%v] said: %v\n", username, ev.Text)

				//only listen if tha masta calls me
				if username == master && strings.Contains(ev.Text, "gobot") {
					replyText := ""

					if strings.Contains(ev.Text, "deploy") && strings.Contains(ev.Text, "kubernetes") && kubeScript != "" {
						replyText = "Building your kubernetes now"
						rtm.SendMessage(rtm.NewOutgoingMessage(replyText, ev.Channel))
						postToUi("One kubernetes cluster coming up!")
						wg.Add(1)

						go func() {
							output, err := exec.Command(kubeScript).CombinedOutput()
							if err != nil {
								rtm.SendMessage(rtm.NewOutgoingMessage(fmt.Sprintf("Ouch, creating the kubernetes cluster failed. Error: %v", err), ev.Channel))
							}
							fmt.Printf("kubeScript output: %s", output)
							rtm.SendMessage(rtm.NewOutgoingMessage(fmt.Sprintf("Kube created: %s ", output), ev.Channel))

							postToUi("Kubernetes cluster created!")
							postToUi("__roll")

							wg.Done()
						}()

					}

					if strings.Contains(ev.Text, "deploy") && strings.Contains(ev.Text, "minecraft") && mineScript != "" {
						wg.Wait()
						replyText = "One minecraft coming up"
						rtm.SendMessage(rtm.NewOutgoingMessage(replyText, ev.Channel))
						postToUi(replyText)

						go func() {
							output, err := exec.Command(mineScript).CombinedOutput()
							if err != nil {
								rtm.SendMessage(rtm.NewOutgoingMessage(fmt.Sprintf("Ouch, creating the minecraft server failed. Error: %v", err), ev.Channel))
								postToUi(fmt.Sprintf("Ouch, creating the minecraft server failed. Error: %v", err))
							}
							fmt.Printf("mineScript output: %s", output)
							rtm.SendMessage(rtm.NewOutgoingMessage(fmt.Sprintf("Minecraft server created: %s ", output), ev.Channel))

							postToUi("Minecraft server created!")
							postToUi("__roll")

							output, err = exec.Command(mineIpScript).CombinedOutput()
							if err != nil {
								rtm.SendMessage(rtm.NewOutgoingMessage(fmt.Sprintf("Getting IP failed. Error: %v", err), ev.Channel))
							}

							postToUi(fmt.Sprintf("Minecraft IP: %s", output))
							rtm.SendMessage(rtm.NewOutgoingMessage(fmt.Sprintf("Minecraft IP: %s", output), ev.Channel))

						}()

					}

					if replyText == "" {
						rtm.SendMessage(rtm.NewOutgoingMessage("I don't understand", ev.Channel))
					}
				}

				if username == "martin" {
					rtm.SendMessage(rtm.NewOutgoingMessage("Martin shutup", ev.Channel))
				}

			case *slack.UserTypingEvent:
				userinfo, err := api.GetUserInfo(ev.User)
				if err != nil {
					return
				}
				fmt.Printf("User typing: %s\n", userinfo.Name)

			case *slack.RTMError:
				fmt.Printf("Error: %s\n", ev.Error())

			case *slack.InvalidAuthEvent:
				fmt.Printf("Invalid credentials")
				return

			default:
			}
		}()
	}
}
