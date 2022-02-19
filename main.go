package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())  // seeds for a random number
	var usedAdventures [5]int         // int array which is checked to prevent duplicate adventures
	var chances = 1 + rand.Intn(50-1) // random number from 1-50, used for roulette
	var void string                   // basically /dev/null
	var birdName string               // the name you set for your birdie
	var input string                  // the input for the bird shell
	var birdHealth int = 100          // bird health
	var birdKarma int = 100           // bird karma
	var birdArtifacts int             // bird artifacts
	var used string = "This adventure has already been used, reroll!"

	fmt.Println("Welcome to Librebird.")
	fmt.Println("This amazing piece of software has been Created by the almighty Coolbird.")
	fmt.Println("I am not responsible for any heart related issue you may get while playing this game")
	fmt.Println("Do you understand these terms? (y/n)")
	fmt.Scanln(&void) // i dont care if you understand or not

	fmt.Println("You are a birdie. You are green. You enjoy eating plants. You like trees. What will be your name?")
	fmt.Println("(No spaces!!!)")
	fmt.Scanln(&birdName)

	if birdName == "" || birdName == " " {

		fmt.Println("Idiot. That is not a valid name. Think about what you have done.")
		fmt.Println("Come back when you are ready.")
		time.Sleep(2 * time.Second)
		fmt.Println("Press ENTER to continue")
		fmt.Println("\033[8m")
		fmt.Scanln()
		fmt.Println("\033[28m")
		os.Exit(0)

	}

	fmt.Println("Wow. " + birdName + " Is such a beautiful name!")
	time.Sleep(2 * time.Second)
	fmt.Print("\033[H\033[2J")
	fmt.Println("Initializing Bird Boot sequence...")
	time.Sleep(2 * time.Second)
	fmt.Println("Starting brain...")
	time.Sleep(3 * time.Second)
	fmt.Println("Readying the wings...")
	time.Sleep(4 * time.Second)
	fmt.Println("Bird initalized. Booting into bird shell")
	time.Sleep(2 * time.Second)
	fmt.Print("\033[H\033[2J")

	fmt.Println("Welcome to BirdOS ver 1.0")
	time.Sleep(1 * time.Second)
	fmt.Println("From here you will be command this bird to victory and success. Good luck!")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("We woo!!! you have a new message! type 'message' to view this message!!!")
	time.Sleep(1 * time.Second)

	for { // infinite loop

		if birdHealth <= 0 {

			fmt.Print("\033[H\033[2J")
			fmt.Println("You have failed in saving your bird. This bird is now dead. Thanks to you, you have murdered this bird, who did nothing wrong.")
			time.Sleep(2 * time.Second)
			fmt.Println("It is time for this bird to rest.")
			time.Sleep(2 * time.Second)
			fmt.Println("Press ENTER to quit the game, knowing you have killed an innocent bird")
			fmt.Println("\033[8m")
			fmt.Scanln(&void)
			fmt.Println("\033[28m")
			os.Exit(0)
		}

		if birdArtifacts > 3 {
			fmt.Print("\033[H\033[2J")
			fmt.Println("Congratulations, you have won the bird game.")
			time.Sleep(3 * time.Second)
			fmt.Println("You have now mastered the art of commanding a Bird.")
			time.Sleep(3 * time.Second)
			fmt.Println("You now have a certificate of bird mastery signed by me.")
			time.Sleep(3 * time.Second)
			fmt.Println("Press ENTER, knowing you are victorious.")
			fmt.Println("\033[8m")
			fmt.Scanln(&void)
			fmt.Println("\033[28m")
			os.Exit(0)
		}

		fmt.Print("[" + birdName + "@bird]: ")
		fmt.Scanln(&input)

		switch input {
		case "ver":
			fmt.Println("BirdOS version: 1.0")
			fmt.Println("BirdOS developed by GoBird Enterprises")

		case "message":
			fmt.Println("From: bird birdinson")
			fmt.Println("To: " + birdName)
			fmt.Println("")
			fmt.Println("hi. you are now controlling a bird. you need to know how to control bird or you lose. birds are hard creature to learn.")
			fmt.Println("use this nice command provided by bird os. it is caled 'help'. it include commandes to help you and bird")
			fmt.Println("you will need to survive. this only possible if you strong enough. good luck")

		case "health":
			fmt.Println("Your bird's health is: ", birdHealth)

		case "help":
			fmt.Println("To win you must collect 4 artifacts. To do this you must either use `adventures` or `roulette`")
			fmt.Println("list of all command:")
			fmt.Println("")
			fmt.Println("ver - print out version of birdos")
			fmt.Println("message - shows ur new messages")
			fmt.Println("health - shows ur bird health")
			fmt.Println("artifacts - shows your artifacts")
			fmt.Println("karma - shows your karma")
			fmt.Println("clear- clears the screen")
			fmt.Println("there r more but i cant be bothered to list them all. have fun")

		case "clear":
			fmt.Print("\033[H\033[2J")

		case "karma":
			fmt.Println("Your karma is:", birdKarma)

		case "artifacts":
			fmt.Println("You have:", birdArtifacts, "artifacts.")

		case "roulette":
			var terms string
			fmt.Println("Welcome to Bird Roulette!")
			fmt.Println("You guess if a random number from 1 to 100 is lower or higher than 50!")
			fmt.Println("If you guess right, you get an artifact, if you guess wrong, you die.")
			fmt.Println("Do you understand these terms? (y/n)")
			fmt.Scanln(&terms)

			switch terms {

			case "y":
				var guess string
				fmt.Println("Input your guess (lower/higher)")
				fmt.Scanln(&guess)

				switch guess {
				case "lower":
					if chances < 50 {
						fmt.Println("Congratualations. You have guessed correctly. You win an artifact.")
						birdArtifacts++
						fmt.Println("You now have: ", birdArtifacts, " artifacts")

					} else {
						fmt.Println("You have guessed wrongly.")
						fmt.Println("Goodbye!")
						time.Sleep(2 * time.Second)
						birdHealth = 0
					}
				case "higher":
					if chances > 50 {
						fmt.Println("Congratualations. You have guessed correctly. You win an artifact.")
						birdArtifacts++
						fmt.Println("You now have: ", birdArtifacts, " artifacts")

					} else {
						fmt.Println("You have guessed wrongly.")
						fmt.Println("Goodbye!")
						time.Sleep(2 * time.Second)
						birdHealth = 0

					}
				default:
					fmt.Println("Invalid.")
				}
			case "n":
				fmt.Println("Then why did you even come here??")

			default:
				fmt.Println("Learn to read next time..")
			}

		case "whoami":
			fmt.Println(birdName)

		case "cd":
			fmt.Println("Command 'cd' is not currently available. You can buy it at librebird.com for 19,99!")

		case "ls":
			fmt.Println("brain")
			fmt.Println("boot")
			fmt.Println("dev")
			fmt.Println("boot.conf")

		case "pwd":
			fmt.Println("/")

		case "adventures":
			var rand = 1 + rand.Intn(4-1) // random number from 1 to the amount of adventures

			var decision string

			switch rand {
			case 1:

				if usedAdventures[1] == 1 {

					fmt.Println(used)
					break
				}

				fmt.Println("You have found a ball of seeds. Would you indulge in eating these seeds? (y/n)")
				fmt.Scanln(&decision)

				switch decision {
				case "y":
					fmt.Println("You have decided you will eat the seeds. They were quite nice. But you feel sick.")
					fmt.Println("These seeds have been poisoned with Fluoride. Your death is inevitable.")
					fmt.Println("Goodbye, World! ", rand)
					time.Sleep(6 * time.Second)
					birdHealth = 0
				case "n":
					fmt.Println("You have decided not to eat the ball of seeds.")

				default:
					fmt.Println("Idiot. That is not a valid option.")
				}
				usedAdventures[1] = 1
			case 2:

				if usedAdventures[2] == 1 {

					fmt.Println(used)
					break
				}

				fmt.Println("You decided you want to go to Africa.")
				fmt.Println("You dont know what to bring with you, a chess board, a fish or seeds")
				fmt.Println("(chess/fish/seeds)")
				fmt.Scanln(&decision)

				switch decision {

				case "chess":

					fmt.Println("You have brung a chess board and pieces to Africa.")
					fmt.Println("All of your bird friends laughed at you.")
					time.Sleep(4 * time.Second)
					fmt.Println("You are crossing the Meditteranian until you see a Falcon. He is very angry and wants to murder you.")
					time.Sleep(4 * time.Second)
					fmt.Println("Luckily, you have chess pieces, so you throw a pawn in his mouth and he suffocates. You lose 30 karma.")
					birdKarma = birdKarma - 30
					time.Sleep(4 * time.Second)
					fmt.Println("You have reached the Saharan desert.")
					fmt.Println("You see pyramids in the distance, but you also see an oasis. Shall you go to the oasis or the pyramids? (oasis/pyramids)")
					fmt.Scanln(&decision)

					switch decision {

					case "pyramids":

						fmt.Println("You have decided to go to the pyramids. They are very big.")
						fmt.Println("You have found an artifact. It seems shiny. ")
						birdArtifacts++
						time.Sleep(4 * time.Second)
						fmt.Println("You now have: ", birdArtifacts, "artifacts")
						time.Sleep(2 * time.Second)
						fmt.Println("You are dying of thirst.")
						fmt.Println("Do you attempt to reach the river or drink your own piss? (river/piss)")
						fmt.Scanln(&void) // it doesn't matter
						fmt.Println("You have survived.")

					case "oasis":

						fmt.Println("You drink the oasis water")
						time.Sleep(2 * time.Second)
						fmt.Println("...")
						time.Sleep(3 * time.Second)
						fmt.Println("...")
						time.Sleep(3 * time.Second)
						fmt.Println("!!!")
						time.Sleep(2 * time.Second)
						fmt.Println("The water is actually radioactive. You now glow in the dark. Your death is inevitable.")
						fmt.Println("You have 10 seconds to think about what you have done.")
						time.Sleep(10 * time.Second)
						birdHealth = 0

					default:

						fmt.Println("Idiot! that is not a valid option")

					}

				case "fish":
					fmt.Println("You are flying over the meditteranian.")
					time.Sleep(2 * time.Second)
					fmt.Println("The fish is too heavy for you. You are now 30 meters below the sea surface. Your death is inevitable.")
					fmt.Println("You have failed.")
					time.Sleep(4 * time.Second)
					birdHealth = 0

				case "seeds":
					fmt.Println("You are smart and bring seeds to your journey.")
					time.Sleep(4 * time.Second)
					fmt.Println("You are flying over the meditteranian")
					time.Sleep(4 * time.Second)
					fmt.Println("You encounter a large, winged cow. It is angry at you. It will want to murder you.")
					time.Sleep(4 * time.Second)
					fmt.Println("You throw seeds at the cow, but it doesn't do anything.")
					time.Sleep(4 * time.Second)
					fmt.Println("The cow attacks you, damaging your feathers.")
					birdHealth = birdHealth - 40
					fmt.Println("Your health is now: ", birdHealth)
					fmt.Println("You want to escape the cow. You can go underwater or you can destroy the cow with your beak (water/destroy)")
					fmt.Scanln(&decision)
					switch decision {

					case "water":

						fmt.Println("You go underwater...")
						time.Sleep(4 * time.Second)
						fmt.Println("You realize you can't swim. Your death is inevitable.")
						time.Sleep(4 * time.Second)
						birdHealth = 0

					case "destroy":

						fmt.Println("You fight the cow..")
						time.Sleep(4 * time.Second)
						fmt.Println("You smashed your beak into the cow, It exploded into 100 different shards.")
						birdArtifacts++
						fmt.Println("You found an artifact. You now have: ", birdArtifacts, " artifacts.")
						fmt.Println("Continue to Africa or go home? (africa/home)")
						fmt.Scanln(&decision)

						switch decision {

						case "africa":

							fmt.Println("You land in Africa...")
							time.Sleep(4 * time.Second)
							fmt.Println("You get eaten by a Penguin.")
							time.Sleep(4 * time.Second)
							birdHealth = 0

						case "home":

							fmt.Println("You go home.")
							time.Sleep(2 * time.Second)

						default:

							fmt.Println("Wrong option error errorr aahhahghh!!")
						}

					default:

						fmt.Println("Idiot!! Wrong option this is not correct!")

					}

				default:

					fmt.Println("Idiot!! Wrong option this is not correct!")

				}
				usedAdventures[2] = 1
			case 3:

				if usedAdventures[3] == 1 {

					fmt.Println(used)
					break
				}

				fmt.Println("You have flown to New York.")
				fmt.Println("The gas emissions are killing you.")
				fmt.Println("Do you attempt to escape or go in the sewers? (escape/sewers)")
				fmt.Scanln(&decision)

				switch decision {

				case "escape":
					fmt.Println("You are attempting to escape...")
					time.Sleep(2 * time.Second)
					fmt.Println("Do you attempt to fly out or go into a van? (fly/van)")
					fmt.Scanln(&decision)

					switch decision {
					case "van":
						fmt.Println("You have entered a van.")
						fmt.Println("...")
						time.Sleep(3 * time.Second)
						fmt.Println("After 3 days of driving, the door finally opens")
						fmt.Println("You are now in the middle of the Mexican desert.")
						time.Sleep(3 * time.Second)
						fmt.Println("A cactus falls on you and you die.")
						time.Sleep(3 * time.Second)
						birdHealth = 0

					case "fly":
						fmt.Println("You decide to fly out...")
						time.Sleep(3 * time.Second)
						fmt.Println("You are now on top of a powerline.")
						fmt.Println("A man has come up to you,")
						fmt.Println("he says that you need to pick between two programming languages.")
						fmt.Println("C or Javascript, picking the wrong one will result in your death.")
						fmt.Println("Pick one. (C/js)")
						fmt.Scanln(&decision)

						switch decision {

						case "C":
							fmt.Println("...")
							time.Sleep(3 * time.Second)
							fmt.Println("Segmentation fault (core dumped)")
							time.Sleep(3 * time.Second)
							fmt.Println("The man has awarded you an artifact for your correct answer.")
							birdArtifacts++
							fmt.Println("You now have: ", birdArtifacts, " artifacts")

						case "js":
							fmt.Println("...")
							time.Sleep(3 * time.Second)
							fmt.Println("He laughed at you, loaded his pistol and shot you in the head.")
							time.Sleep(3 * time.Second)
							birdHealth = 0

						default:
							fmt.Println("Your answer didn't satisfy the man.")
							fmt.Println("You have been shot in the head.")
							time.Sleep(3 * time.Second)
							birdHealth = 0
						}
					default:
						fmt.Println("Invalid answer. bud.")
						time.Sleep(3 * time.Second)
						fmt.Println("Just because of that, I am going to murder you.")
						time.Sleep(3 * time.Second)
						fmt.Println("Have fun, kid.")
						time.Sleep(3 * time.Second)
						birdHealth = 0
					}
				case "sewers":
					fmt.Println("You have decided to go through the sewers.")
					time.Sleep(3 * time.Second)
					fmt.Println("You meet a mutant ninja turtle.")
					fmt.Println("He gives you an artifact.")
					time.Sleep(2 * time.Second)
					fmt.Println("But it seems... fishy...")
					time.Sleep(2 * time.Second)
					fmt.Println("*boom bamm kapoof kapow boom boom!!*")
					time.Sleep(2 * time.Second)
					fmt.Println("The artifact exploded in your arms. It appears that there was some sort of mutant virus in the artifact..")
					fmt.Println("You are now becoming mutant. Your brain death is inevitable.")
					time.Sleep(6 * time.Second)
					birdHealth = 0

				default:
					fmt.Println("Can you read? do you see the valid options, mister or madam or maam?")
				}
				usedAdventures[3] = 1
			case 4:

				if usedAdventures[4] == 1 {

					fmt.Println(used)
					break
				}

				fmt.Println("You decide to play chess with Grandmaster Hikaru Nakamura")
				fmt.Println("He is about to checkmate you, what will you do?")
				fmt.Println("Move your pawn or move your queen (pawn/queen)")
				fmt.Scanln(&decision)

				switch decision {

				case "pawn":
					fmt.Println("You have decided to move your pawn.")
					fmt.Println("It was completely useless. You have been checkmated.")
					time.Sleep(2 * time.Second)
					fmt.Println("You are now sentenced to death.")
					time.Sleep(2 * time.Second)
					fmt.Println("Goodbye.")
					time.Sleep(2 * time.Second)
					birdHealth = 0

				case "queen":
					fmt.Println("You move the queen.")
					time.Sleep(2 * time.Second)
					fmt.Println("Its a miracle! You managed to mate in 1 him!")
					time.Sleep(2 * time.Second)
					fmt.Println("After being defeated by a bird, Hikaru Nakamura lives in self isolation, stopped playing chess and will live in humiliation for the rest of his life.")
					time.Sleep(2 * time.Second)
					fmt.Println("You have gained an artifact but lost 20 karma.")
					birdArtifacts++
					birdKarma = birdKarma - 20
					fmt.Println("You now have:", birdArtifacts, "artifacts")
					fmt.Println("You now have: ", birdKarma, "karma")

				default:
					fmt.Println("After failing to decide, Hikaru checkmates you.")
					fmt.Println("You are now going to die, cya")
					time.Sleep(2 * time.Second)
					birdHealth = 0

				}
				usedAdventures[4] = 1
			}

		default:
			fmt.Println("Bird does not understand command '" + input + "', type 'help' for bird understandable commands.")

		}

	}

}
