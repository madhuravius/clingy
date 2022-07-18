package internal

import (
	"embed"
	"html/template"
	"log"
	"os"

	"clingy/lib"
)

//go:embed templates
var templates embed.FS

// GenerateHTMLReport - generate a HTML report by reading from a file and writing to a specified destination
func GenerateHTMLReport(logger *log.Logger, data *lib.ClingyTemplate, outFile string) error {
	logger.Println("Generating HTML report", data.Label, outFile)

	inputTemplateBytes, err := templates.ReadFile("templates/simple-report.template.html")
	if err != nil {
		logger.Println("Error in reading embed template", err)
		return err
	}

	inputTemplate, err := template.New("report").Parse(string(inputTemplateBytes))
	if err != nil {
		logger.Println("Error in creating template from templatedir", err)
		return err
	}

	out, err := os.Create(outFile)
	if err != nil {
		logger.Println("Error in creating outfile", outFile, err)
		return err
	}
	defer func(out *os.File) {
		err = out.Close()
		if err != nil {
			logger.Println("Error in closing outFile", outFile, err)
			// unable to safely return in closure/defer func
		}
	}(out)

	if err = inputTemplate.ExecuteTemplate(out, "report", data); err != nil {
		logger.Println("Error in executing on template", err)
		return err
	}

	logger.Println("Completed generating HTML report", data.Label, outFile)
	return nil
}
