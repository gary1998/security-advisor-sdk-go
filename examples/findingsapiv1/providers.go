package examples

import (
	"fmt"

	"github.com/IBM/go-sdk-core/v3/core"
	"github.com/ibm-cloud-security/security-advisor-sdk-go/v3/findingsapiv1"
)

//ListProviders Lists all providers under a account
func ListProviders() {

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	authenticator := &core.IamAuthenticator{
		ApiKey: apiKey,
		URL:    url, //use for dev/preprod env
	}
	service, _ := findingsapiv1.NewFindingsApiV1(&findingsapiv1.FindingsApiV1Options{
		Authenticator: authenticator,
		URL:           "https://us-south.secadvisor.cloud.ibm.com/findings", //Specify url or use default
	})

	var listProvidersOptions = service.NewListProvidersOptions(accountID)
	listProvidersOptions.SetLimit(5)
	listProvidersOptions.SetStartProviderID("a")
	listProvidersOptions.SetEndProviderID("p")

	res, _, err := service.ListProviders(listProvidersOptions)
	if err != nil {
		fmt.Println("Failed to get list of providers: ", err)
	} else {
		fmt.Printf(`Found %d Providers between "a" and "p". Limit is set to 5 per page.`, len(res.Providers))
		fmt.Println()
		if len(res.Providers) > 0 {
			fmt.Println("Providers 1 id: ", *res.Providers[0].ID)
			fmt.Println("Providers 1 name: ", *res.Providers[0].Name)

		}
	}

}
