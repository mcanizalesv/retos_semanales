package main

import (
	"challenge5/files"
	"fmt"
	"log"
	"math"
	"strings"
)

func main() {
	type optionUsers int
	const (
		InvalidOption optionUsers = iota
		Register 
		Login
	)

	// Archivo donde se guardarán los registros
	const recordsFile = "records.json" //Buscar lo que son las variables de entorno. 

	// Cargar registros desde el archivo y los guarda en el map "records"
	records := files.LoadRecords(recordsFile)

	// Mensaje de bienvenida
	fmt.Println("This is Spaced repetition app\nIt helps you to learn\nusing the spaced repetition algorithm!!!\nGood luck!")
	
	var OptionUsers optionUsers
	var NameInput string
	var InfoUser files.User
	var WordApparition int

	// Selección de la opción por parte del usuario
	fmt.Println("Did you want to (into the num, for example, 1 or 2):\n0.Exit\n1.Register\n2.Login")
	_, err := fmt.Scan(&OptionUsers)
	if err != nil {
		OptionUsers = InvalidOption
	}

	switch OptionUsers {

	case InvalidOption:

		fmt.Println("Thanks for using this app")

	case Register:

		// El usuario ingresa el nombre
		fmt.Println("Please tell me your name?")
		_, err := fmt.Scan(&NameInput) //Traté hacerlo con Scanln, pero se corta el flujo
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		NameInput = strings.ReplaceAll(NameInput, " ", "")
		NameInput = strings.ToLower(NameInput)
		fmt.Println("The name is",NameInput)
		

		// Validamos que no exista otro registro con el mismo nombre.
		for {
			if _, exists := records[NameInput]; exists {
				fmt.Printf("This name %s already exists. Please, enter a different name: ", NameInput)
				_, err := fmt.Scan(&NameInput)
				if err != nil {
					log.Fatal(err) //Buscar alternativa que no cierre la app
			
				}
			} else {
				break
			}
		}

		// Inicialización de un nuevo usuario
		users := files.User{Words: files.Words{
			Learning:    files.Progress{},
			Walking:     files.Progress{},
			Certain:     files.Progress{},
			Intricate:   files.Progress{},
			Calculation: files.Progress{},
		}}

		// Guarda el nuevo registro en el mapa
		records[NameInput] = users
		InfoUser = records[NameInput]
		// Guarda los registros actualizados en el archivo
		files.SaveRecords(records,recordsFile)

		// Imprime los registros actuales
		fmt.Println("Records updated successfully!")

	case Login:
		fmt.Println("Please enter your name:")
		_, err := fmt.Scan(&NameInput)
		NameInput = strings.ToLower(NameInput)
		NameInput = strings.ReplaceAll(NameInput, " ", "")
		
		if err != nil {
			log.Fatal(err)
		}

		// Verificar si el usuario existe en los registros
		if user, exists := records[NameInput]; exists {
			fmt.Println("User logged in successfully!")
			InfoUser = user //Guardamos el usuario logeado en la variable InfoUser para luego ser usada
			fmt.Println("")
			
			//****************** Inicia menu learn and progress**********************
			selectedMenuLearn := 1

			round := 0					
			wordsEnglish := []string{"learning", "walking", "certain", "intricate", "calculation"}
			englishSpanish := map[string]string{"learning": "aprender", "walking": "caminar", "certain": "cierto", "intricate": "intricado", "calculation": "calculo"}
		
		
			correctEnglishSpanish := map[string]bool{"learning": false, "walking": false, "certain": false, "intricate": false, "calculation": false}
		
			var greatBadWordInput string
			
			
			for selectedMenuLearn == 1 {
				fmt.Println(greatBadWordInput + " What did you want to do?")
				fmt.Println("1. Go to the next word")	
				fmt.Println("2. See my progress")	
				fmt.Println("3. Finish")	
		
				round++ // para ver nro de round
				fmt.Scanln(&selectedMenuLearn)
				fmt.Println("")
				if selectedMenuLearn == 1 {
					fmt.Println("round:", round ," selectedMenuLearn:", selectedMenuLearn)
				}
								
				switch selectedMenuLearn {			
					case 1:
						i := 0				
						var wordInput string	
						wordAsk := wordsEnglish[i]
						fmt.Printf("What does %s means in Spanish?\n", wordAsk)
						fmt.Scanln(&wordInput)	
						wordSpanish := englishSpanish[wordsEnglish[i]] // aprender
						
						if wordInput != wordSpanish {
							correctEnglishSpanish[wordAsk] = false
							temp := wordAsk
							wordsEnglish[i] = wordsEnglish[i+1]   
							wordsEnglish[i+1] = temp
							greatBadWordInput = "Bad!"
						} else {

							//Actualizamos campo de la palabra que recién aparece
							files.UpdateProgress(&InfoUser.Words, wordAsk, true)
							correctEnglishSpanish[wordAsk] = true	

							// Switch para actualizar solo la palabra actual
							switch wordAsk {
							case "learning":
								WordApparition = int(InfoUser.Words.Learning.Apparitions)
							case "walking":
								WordApparition = int(InfoUser.Words.Walking.Apparitions)
							case "certain":
								WordApparition = int(InfoUser.Words.Certain.Apparitions)
							case "intricate":
								WordApparition = int(InfoUser.Words.Intricate.Apparitions)
							case "calculation":
								WordApparition = int(InfoUser.Words.Calculation.Apparitions)
							}

							//Se actualizan las posiciones del array wordsEnglish
							newPosition := int(math.Pow(2, float64(WordApparition)))
							wordsEnglishcOPY := make([]string, len(wordsEnglish))
							copy(wordsEnglishcOPY, wordsEnglish)
				
							if len(wordsEnglish) > newPosition {
								wordsEnglish[newPosition] = wordsEnglish[0]
								for i := 0; i < newPosition; i++ {
									wordsEnglish[i] = wordsEnglishcOPY[i+1]
								}	
							} else {
								wordsEnglish[len(wordsEnglish)-1] = wordsEnglish[0]
								for i := 0; i < len(wordsEnglish) - 1; i++ {
									wordsEnglish[i] = wordsEnglishcOPY[i+1]
								}	
							}			
							greatBadWordInput = "Great!"	
							// Aquí se guardan los registros actualizados en el json
							records[NameInput] = InfoUser
							files.SaveRecords(records, recordsFile)	
						}	
						fmt.Println("")	
						
						case 2:	
						//muestra el progreso del usuario y vuelve a mostrar el menu
						files.PrintProgress(InfoUser.Words)
						selectedMenuLearn = 1
					default:		
					fmt.Println("Thank you for use this app!")
							
				}			
		
			}
			//****************** Fin menu learn and progress*********************
		} else {
			fmt.Println("User not found.")
		}

	default:
		fmt.Println("Invalid option selected")
	}
	
}