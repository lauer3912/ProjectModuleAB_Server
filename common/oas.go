package common

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/tonychee7000/oas"
)

type OasClient struct {
	*oas.OasClient
}

func (o *OasClient) GetOasVaultId(name string) (string, error) {
	v := new(oas.VaultsList)
	for {
		id, v, err := o.ListVaults(-1, v.Marker)
		beego.Debug("OAS request ID:", id)
		if err != nil {
			return "", err
		}
		beego.Debug(v)
		for _, xv := range v.VaultList {
			beego.Debug(xv.VaultName, name)
			if xv.VaultName == name {
				return xv.VaultID, nil
			}
		}
		if v.Marker == "" {
			return "", fmt.Errorf("Vault not found")
		}
	}
}

func NewOasClient(endpoint string) (*OasClient, error) {
	oasPort := beego.AppConfig.DefaultInt("aliapi::oasport", 80)
	oasUseSSL := beego.AppConfig.DefaultBool("aliapi::oasusessl", false)
	o := new(OasClient)
	o.OasClient = oas.NewOasClient(
		endpoint,
		beego.AppConfig.String("aliapi::apikey"),
		beego.AppConfig.String("aliapi::secret"),
		oasPort,
		oasUseSSL,
	)
	return o, nil
}
