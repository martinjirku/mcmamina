import { basename, join } from "path";

interface Sponsor {
  url?: string;
  img: string;
}
const sponzori: Sponsor[] = [
  {
    url: "http://www.banskabystrica.sk/",
    img: "https://www.mcmamina.sk/nova_stranka/wp-content/uploads/2019/12/mestoBB_s_fin_podporou.jpg",
  },
  {
    url: "https://www.bbsk.sk/",
    img: "https://www.mcmamina.sk/nova_stranka/wp-content/uploads/2019/10/COA-fin-podpora-1.jpg",
  },
  {
    url: "https://www.tvojeharmony.sk/",
    img: "https://www.mcmamina.sk/nova_stranka/wp-content/uploads/2018/10/HARMANEC-1.jpg",
  },
  {
    url: "http://www.mclloyds.eu/",
    img: "https://www.mcmamina.sk/nova_stranka/wp-content/uploads/2018/10/mclloyds.png",
  },
  {
    url: "https://www.bbxnet.sk/",
    img: "https://www.mcmamina.sk/nova_stranka/wp-content/uploads/2018/10/bbxnet.jpg",
  },
  {
    url: "https://www.europasc.sk/",
    img: "https://www.mcmamina.sk/nova_stranka/wp-content/uploads/2018/10/europa.png",
  },
  {
    url: "http://emsbb.sk/",
    img: "https://www.mcmamina.sk/nova_stranka/wp-content/uploads/2018/10/ems-logo-1160x535.jpg",
  },
  {
    url: "https://www.nadaciaspp.sk/",
    img: "https://www.mcmamina.sk/nova_stranka/wp-content/uploads/2018/10/logo-spp.jpg",
  },
  {
    url: "https://tesco.sk/",
    img: "https://www.mcmamina.sk/nova_stranka/wp-content/uploads/2018/10/tesco.png",
  },
  {
    url: "https://lunter.com/",
    img: "https://www.mcmamina.sk/nova_stranka/wp-content/uploads/2018/10/LUNTER_LOGO_RGB-NOVE.png",
  },
  {
    img: "https://www.mcmamina.sk/nova_stranka/wp-content/uploads/2018/10/BOP.jpg",
  },
  {
    url: "https://www.dru.sk/",
    img: "https://www.mcmamina.sk/nova_stranka/wp-content/uploads/2018/10/Logo-DRU.jpg",
  },
  {
    url: "https://www.facebook.com/KatarinaMalovaPhotography/",
    img: "https://www.mcmamina.sk/nova_stranka/wp-content/uploads/2018/12/Katarina_Malova_logo.jpg",
  },
  {
    url: "https://www.cosmopolitanbb.sk/sk/",
    img: "https://www.mcmamina.sk/nova_stranka/wp-content/uploads/2019/02/Cosmopolitan.jpg",
  },
  {
    url: "http://www.bdnr.sk/",
    img: "https://www.mcmamina.sk/nova_stranka/wp-content/uploads/2019/02/BDNR_logo_nove.jpg",
  },
  {
    url: "https://www.marykay.sk/",
    img: "https://www.mcmamina.sk/nova_stranka/wp-content/uploads/2019/02/Mary-Kay.jpg",
  },
  {
    img: "https://www.mcmamina.sk/nova_stranka/wp-content/uploads/2020/02/lukas_obernauer.jpg",
  },
  {
    url: "https://www.expodom.sk/",
    img: "https://www.mcmamina.sk/nova_stranka/wp-content/uploads/2021/03/expodom.png",
  },
  {
    url: "https://www.coop.sk/sk/nadacia-coop-jednota",
    img: "https://www.mcmamina.sk/nova_stranka/wp-content/uploads/2021/03/coop-jednota-1.png",
  },
  {
    url: "https://www.raiffeisen.sk/",
    img: "https://www.mcmamina.sk/nova_stranka/wp-content/uploads/2021/05/reiffeisen-banka.png",
  },
  {
    url: "http://www.panfoto.sk/",
    img: "https://www.mcmamina.sk/nova_stranka/wp-content/uploads/2021/05/panfoto-2.png",
  },
  {
    url: "http://www.daliprint.eu/",
    img: "https://www.mcmamina.sk/nova_stranka/wp-content/uploads/2021/05/daliprint.png",
  },
  {
    img: "https://www.mcmamina.sk/nova_stranka/wp-content/uploads/2021/06/prirodzeny-svet.png",
  },
  {
    url: "https://www.draculagym.sk/",
    img: "https://www.mcmamina.sk/nova_stranka/wp-content/uploads/2021/06/logo-new-2020.png",
  },
  {
    url: "https://www.facebook.com/Alexis-detsky-obchodik-453842445025893/",
    img: "https://www.mcmamina.sk/nova_stranka/wp-content/uploads/2021/06/alexis_logo.png",
  },
  {
    url: "https://www.dedoles.sk/",
    img: "https://www.mcmamina.sk/nova_stranka/wp-content/uploads/2021/06/image002.png",
  },
  {
    url: "http://evijosasova.sk/",
    img: "https://www.mcmamina.sk/nova_stranka/wp-content/uploads/2021/06/Evijo-Sasova.png",
  },
  {
    url: "https://intaxi.sk/",
    img: "https://www.mcmamina.sk/nova_stranka/wp-content/uploads/2021/06/Intaxi.png",
  },
  {
    url: "https://www.japitex.sk/",
    img: "https://www.mcmamina.sk/nova_stranka/wp-content/uploads/2021/06/japitex-logo.png",
  },
  {
    url: "https://www.zvolensky.com/",
    img: "https://www.mcmamina.sk/nova_stranka/wp-content/uploads/2021/06/Logo-CMYK.png",
  },
  {
    url: "https://www.facebook.com/Pono%C5%BEk%C3%A1%C4%8Di-od-Ivu%C5%A1ky-2174200106172832/",
    img: "https://www.mcmamina.sk/nova_stranka/wp-content/uploads/2021/06/Ponozkaci-od-Ivusky.png",
  },
  {
    url: "https://www.bouncepark.sk/",
    img: "https://www.mcmamina.sk/nova_stranka/wp-content/uploads/2021/06/bounce-logo-17-digital-01.png",
  },
  {
    url: "https://donovalkovo.sk/en/",
    img: "https://www.mcmamina.sk/nova_stranka/wp-content/uploads/2021/06/Donovalkovo_logo-01.png",
  },
  {
    img: "https://www.mcmamina.sk/nova_stranka/wp-content/uploads/2021/11/eat.jpg",
  },
  {
    url: "https://www.cinema.sk/",
    img: "https://www.mcmamina.sk/nova_stranka/wp-content/uploads/2021/11/cinema.jpg",
  },
  {
    url: "https://beelol.eu/",
    img: "https://www.mcmamina.sk/nova_stranka/wp-content/uploads/2021/11/beeLOL.jpg",
  },
  {
    url: "https://www.mojadm.sk/",
    img: "https://www.mcmamina.sk/nova_stranka/wp-content/uploads/2021/11/dm.png",
  },
  {
    url: "https://www.alaska.sk/",
    img: "https://www.mcmamina.sk/nova_stranka/wp-content/uploads/2021/11/alaska-food.jpg",
  },
  {
    url: "https://www.mixit.sk/",
    img: "https://www.mcmamina.sk/nova_stranka/wp-content/uploads/2021/11/mixit.jpg",
  },
  {
    url: "https://www.coopka.sk/",
    img: "https://www.mcmamina.sk/nova_stranka/wp-content/uploads/2021/12/coopka.jpg",
  },
  {
    url: "https://www.artforum.sk/clanky/knihkupectva#banska-bystrica",
    img: "https://www.mcmamina.sk/nova_stranka/wp-content/uploads/2022/04/Dizajn-bez-nazvu-2.png",
  },
  {
    url: "https://www.siko.sk/",
    img: "https://www.mcmamina.sk/nova_stranka/wp-content/uploads/2022/04/Dizajn-bez-nazvu.png",
  },
  {
    url: "https://www.martinus.sk/",
    img: "https://www.mcmamina.sk/nova_stranka/wp-content/uploads/2022/04/Dizajn-bez-nazvu-1.png",
  },
];
export const getSponsors = async () => {
  return sponzori.map((s) => {
    const img = s.img ? join("/images/sponzori/", basename(s.img)) : undefined;
    return {
      ...s,
      img,
    };
  });
};
