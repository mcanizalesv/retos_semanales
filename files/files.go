package files

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type Progress struct {
	Apparitions     float64 `json:"apparitions"`
	Correct_answers float64 `json:"correct_answers"`
	Retention_rate  float64 `json:"retention_rate"`
}


type Words struct {
	Learning    Progress `json:"learning"`
	Walking     Progress `json:"walking"`
	Certain     Progress `json:"certain"`
	Intricate   Progress `json:"intricate"`
	Calculation Progress `json:"calculation"`
}

type User struct {
	Words Words `json:"words"`
}

// LoadRecords sirve para cargar los registros desde un archivo
 func LoadRecords(recordsFile string) map[string]User {
	file, err := os.Open(recordsFile)
	if err != nil {
		if os.IsNotExist(err) {
			// Si el archivo no existe, devolvemos un mapa vacío - crear archivo
			return make(map[string]User)
		}

		log.Fatal(err)
	}

	defer file.Close()

	var recordsMap map[string]User

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	//Revisar con return early - refactorización
	if len(fileBytes) == 0{
		return make(map[string]User)
	}

    if err := json.Unmarshal(fileBytes, &recordsMap); err != nil {log.Fatal(err)}
	
	return recordsMap
}


// SaveRecords para guardar los registros en un archivo
func SaveRecords(records map[string]User, recordsFile string) {
	bytes, err := json.MarshalIndent(records, "", "  ")
	if err != nil {
		log.Fatal(fmt.Errorf("marshall records:%w",err)) //Evaluar si está bien que falle toda la app si se hace mal la lectura
	}
	err = os.WriteFile(recordsFile, bytes, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func UpdateProgress(words *Words, word string, correct bool) {
	switch word {
	case "learning":
		words.Learning.Apparitions++
		if correct {
			words.Learning.Correct_answers++
		}
		words.Learning.Retention_rate = words.Learning.Correct_answers / words.Learning.Apparitions
	case "walking":
		words.Walking.Apparitions++
		if correct {
			words.Walking.Correct_answers++
		}
		words.Walking.Retention_rate = words.Walking.Correct_answers / words.Walking.Apparitions
	case "certain":
		words.Certain.Apparitions++
		if correct {
			words.Certain.Correct_answers++
		}
		words.Certain.Retention_rate = words.Certain.Correct_answers / words.Certain.Apparitions
	case "intricate":
		words.Intricate.Apparitions++
		if correct {
			words.Intricate.Correct_answers++
		}
		words.Intricate.Retention_rate = words.Intricate.Correct_answers / words.Intricate.Apparitions
	case "calculation":
		words.Calculation.Apparitions++
		if correct {
			words.Calculation.Correct_answers++
		}
		words.Calculation.Retention_rate = words.Calculation.Correct_answers / words.Calculation.Apparitions
	}
}

func PrintProgress(words Words) {
    fmt.Println("Your current progress is:")
    fmt.Println("Words       | Retention Rate")

    // Formatea cada palabra y su tasa de retención
    fmt.Printf("%s | %.0f%%\n", "learning", words.Learning.Retention_rate * 100)
    fmt.Printf("%s | %.0f%%\n", "walking", words.Walking.Retention_rate * 100)
    fmt.Printf("%s | %.0f%%\n", "certain", words.Certain.Retention_rate * 100)
    fmt.Printf("%s | %.0f%%\n", "intricate", words.Intricate.Retention_rate * 100)
    fmt.Printf("%s | %.0f%%\n", "calculation", words.Calculation.Retention_rate * 100)
}

