package services

import (
	"context"

	"jirku.sk/mcmamina/models"
)

type SponsorService struct{}

func NewSponsorService() *SponsorService {
	return &SponsorService{}
}

func (s *SponsorService) GetSponsors(ctx context.Context) ([]models.Sponsor, error) {
	return []models.Sponsor{
		models.NewSponsor("http://www.banskabystrica.sk/", "mestoBB_s_fin_podporou.jpg"),
		models.NewSponsor("https://www.bbsk.sk/", "COA-fin-podpora-1.jpg"),
		models.NewSponsor("https://www.tvojeharmony.sk/", "HARMANEC-1.jpg"),
		models.NewSponsor("http://www.mclloyds.eu/", "mclloyds.png"),
		models.NewSponsor("https://www.bbxnet.sk/", "bbxnet.jpg"),
		models.NewSponsor("https://www.europasc.sk/", "europa.png"),
		models.NewSponsor("http://emsbb.sk/", "ems-logo-1160x535.jpg"),
		models.NewSponsor("https://www.nadaciaspp.sk/", "logo-spp.jpg"),
		models.NewSponsor("https://tesco.sk/", "tesco.png"),
		models.NewSponsor("https://lunter.com/", "LUNTER_LOGO_RGB-NOVE.png"),
		models.NewSponsor("", "BOP.jpg"),
		models.NewSponsor("https://www.dru.sk/", "Logo-DRU.jpg"),
		models.NewSponsor("https://www.facebook.com/KatarinaMalovaPhotography/", "Katarina_Malova_logo.jpg"),
		models.NewSponsor("https://www.cosmopolitanbb.sk/sk/", "Cosmopolitan.jpg"),
		models.NewSponsor("http://www.bdnr.sk/", "BDNR_logo_nove.jpg"),
		models.NewSponsor("https://www.marykay.sk/", "Mary-Kay.jpg"),
		models.NewSponsor("", "lukas_obernauer.jpg"),
		models.NewSponsor("https://www.expodom.sk/", "expodom.png"),
		models.NewSponsor("https://www.coop.sk/sk/nadacia-coop-jednota", "coop-jednota-1.png"),
		models.NewSponsor("https://www.raiffeisen.sk/", "reiffeisen-banka.png"),
		models.NewSponsor("http://www.panfoto.sk/", "panfoto-2.png"),
		models.NewSponsor("http://www.daliprint.eu/", "daliprint.png"),
		models.NewSponsor("", "prirodzeny-svet.png"),
		models.NewSponsor("https://www.draculagym.sk/", "logo-new-2020.png"),
		models.NewSponsor("https://www.facebook.com/Alexis-detsky-obchodik-453842445025893/", "alexis_logo.png"),
		models.NewSponsor("https://www.dedoles.sk/", "image002.png"),
		models.NewSponsor("http://evijosasova.sk/", "Evijo-Sasova.png"),
		models.NewSponsor("https://intaxi.sk/", "Intaxi.png"),
		models.NewSponsor("https://www.japitex.sk/", "japitex-logo.png"),
		models.NewSponsor("https://www.zvolensky.com/", "Logo-CMYK.png"),
		models.NewSponsor("https://www.facebook.com/Pono%C5%BEk%C3%A1%C4%8Di-od-Ivu%C5%A1ky-2174200106172832/", "Ponozkaci-od-Ivusky.png"),
		models.NewSponsor("https://www.bouncepark.sk/", "bounce-logo-17-digital-01.png"),
		models.NewSponsor("https://donovalkovo.sk/en/", "Donovalkovo_logo-01.png"),
		models.NewSponsor("", "eat.jpg"),
		models.NewSponsor("https://www.cinema.sk/", "cinema.jpg"),
		models.NewSponsor("https://beelol.eu/", "beeLOL.jpg"),
		models.NewSponsor("https://www.mojadm.sk/", "dm.png"),
		models.NewSponsor("https://www.alaska.sk/", "alaska-food.jpg"),
		models.NewSponsor("https://www.mixit.sk/", "mixit.jpg"),
		models.NewSponsor("https://www.coopka.sk/", "coopka.jpg"),
		models.NewSponsor("https://www.artforum.sk/clanky/knihkupectva#banska-bystrica", "Dizajn-bez-nazvu-2.png"),
		models.NewSponsor("https://www.siko.sk/", "Dizajn-bez-nazvu.png"),
		models.NewSponsor("https://www.martinus.sk/", "Dizajn-bez-nazvu-1.png"),
	}, nil
}
