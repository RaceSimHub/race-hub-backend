package iracing

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type IRacing struct {
	Email    string
	Password string
}

func (ir IRacing) GetIRatings(custIds []int) map[int]IRatingData {
	cookies, err := ir.authenticateIracing()
	if err != nil {
		fmt.Println("❌ Erro ao autenticar no iRacing.")
		return nil
	}

	iratingsMap := make(map[int]IRatingData)
	batchSize := 25

	for i := 0; i < len(custIds); i += batchSize {
		batch := custIds[i:min(i+batchSize, len(custIds))]

		// Transform the batch of int into a string
		// Example: [1, 2, 3] -> "1,2,3"
		var batchStr []string
		for _, id := range batch {
			batchStr = append(batchStr, fmt.Sprintf("%v", id))
		}

		url := fmt.Sprintf("https://members-ng.iracing.com/data/member/get?cust_ids=%v&include_licenses=true", strings.Join(batchStr, ","))

		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Add("Cookie", cookies)

		resp, err := http.DefaultClient.Do(req)
		if err != nil || resp.StatusCode != 200 {
			fmt.Printf("❌ Erro ao buscar dados para IDs: %v\n", batch)
			continue
		}

		bodyMember, _ := io.ReadAll(resp.Body)
		resp.Body.Close()

		var memberData MemberDataLink
		if err = json.Unmarshal(bodyMember, &memberData); err != nil {
			fmt.Printf("❌ Erro ao decodificar dados para IDs: %v\n", batch)
			continue
		}

		req, _ = http.NewRequest("GET", memberData.Link, nil)
		req.Header.Add("Cookie", cookies)

		resp, err = http.DefaultClient.Do(req)
		if err != nil || resp.StatusCode != 200 {
			fmt.Printf("❌ Erro ao buscar dados para IDs: %v\n", batch)
			continue
		}

		var data MemberData
		json.NewDecoder(resp.Body).Decode(&data)
		resp.Body.Close()

		for _, member := range data.Members {
			irating := IRatingData{}
			for _, license := range member.Licenses {
				switch {
				case strings.Contains(license.Category, "sports_car"):
					irating.SportsCar = license.IRating
				case strings.Contains(license.Category, "formula_car"):
					irating.FormulaCar = license.IRating
				case strings.Contains(license.Category, "oval"):
					irating.Oval = license.IRating
				}
			}
			iratingsMap[member.CustID] = irating
		}
	}
	return iratingsMap
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (ir IRacing) authenticateIracing() (cookies string, err error) {
	url := "https://members-ng.iracing.com/auth"
	payload := fmt.Sprintf(`{"email":"%s","password":"%s"}`, ir.Email, hashIracingPassword(ir.Email, ir.Password))
	resp, err := http.Post(url, "application/json", strings.NewReader(payload))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var authData AuthResponse
	if err = json.Unmarshal(body, &authData); err != nil {
		return
	}

	headerCookies := resp.Header["Set-Cookie"]
	return strings.Join(headerCookies, "; "), nil
}

func hashIracingPassword(email, password string) string {
	saltedPassword := password + email
	hash := sha256.Sum256([]byte(saltedPassword))

	return base64.StdEncoding.EncodeToString(hash[:])
}
