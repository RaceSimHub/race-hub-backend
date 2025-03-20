package localization

import (
	"embed"
	"encoding/json"
	"fmt"
)

//go:embed *.json
var localeFiles embed.FS

func LoadTranslations(lang string) (map[string]string, error) {
	filePath := fmt.Sprintf("%s.json", lang)

	data, err := localeFiles.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("erro ao carregar arquivo de localização '%s': %w", filePath, err)
	}

	var translations map[string]string
	if err := json.Unmarshal(data, &translations); err != nil {
		return nil, fmt.Errorf("erro ao decodificar as traduções de '%s': %w", filePath, err)
	}

	return translations, nil
}
