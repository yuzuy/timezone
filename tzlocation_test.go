package tzlocation_test

import (
	"testing"

	"github.com/yuzuy/tzlocation"
)

var data = []tzlocation.TZLocation{
	tzlocation.AfricaAbidjan,
	tzlocation.AfricaAccra,
	tzlocation.AfricaAddisAbaba,
	tzlocation.AfricaAlgiers,
	tzlocation.AfricaAsmara,
	tzlocation.AfricaBamako,
	tzlocation.AfricaBangui,
	tzlocation.AfricaBanjul,
	tzlocation.AfricaBissau,
	tzlocation.AfricaBlantyre,
	tzlocation.AfricaBrazzaville,
	tzlocation.AfricaBujumbura,
	tzlocation.AfricaCairo,
	tzlocation.AfricaCasablanca,
	tzlocation.AfricaCeuta,
	tzlocation.AfricaConakry,
	tzlocation.AfricaDakar,
	tzlocation.AfricaDaresSalaam,
	tzlocation.AfricaDjibouti,
	tzlocation.AfricaDouala,
	tzlocation.AfricaElAaiun,
	tzlocation.AfricaFreetown,
	tzlocation.AfricaGaborone,
	tzlocation.AfricaHarare,
	tzlocation.AfricaJohannesburg,
	tzlocation.AfricaJuba,
	tzlocation.AfricaKampala,
	tzlocation.AfricaKhartoum,
	tzlocation.AfricaKigali,
	tzlocation.AfricaKinshasa,
	tzlocation.AfricaLagos,
	tzlocation.AfricaLibreville,
	tzlocation.AfricaLome,
	tzlocation.AfricaLuanda,
	tzlocation.AfricaLubumbashi,
	tzlocation.AfricaLusaka,
	tzlocation.AfricaMalabo,
	tzlocation.AfricaMaputo,
	tzlocation.AfricaMaseru,
	tzlocation.AfricaMbabane,
	tzlocation.AfricaMogadishu,
	tzlocation.AfricaMonrovia,
	tzlocation.AfricaNairobi,
	tzlocation.AfricaNdjamena,
	tzlocation.AfricaNiamey,
	tzlocation.AfricaNouakchott,
	tzlocation.AfricaOuagadougou,
	tzlocation.AfricaPortoNovo,
	tzlocation.AfricaSaoTome,
	tzlocation.AfricaTripoli,
	tzlocation.AfricaTunis,
	tzlocation.AfricaWindhoek,
	tzlocation.AmericaAdak,
	tzlocation.AmericaAnchorage,
	tzlocation.AmericaAnguilla,
	tzlocation.AmericaAntigua,
	tzlocation.AmericaAraguaina,
	tzlocation.AmericaArgentinaBuenosAires,
	tzlocation.AmericaArgentinaCatamarca,
	tzlocation.AmericaArgentinaCordoba,
	tzlocation.AmericaArgentinaJujuy,
	tzlocation.AmericaArgentinaLaRioja,
	tzlocation.AmericaArgentinaMendoza,
	tzlocation.AmericaArgentinaRioGallegos,
	tzlocation.AmericaArgentinaSalta,
	tzlocation.AmericaArgentinaSanJuan,
	tzlocation.AmericaArgentinaSanLuis,
	tzlocation.AmericaArgentinaTucuman,
	tzlocation.AmericaArgentinaUshuaia,
	tzlocation.AmericaAruba,
	tzlocation.AmericaAsuncion,
	tzlocation.AmericaAtikokan,
	tzlocation.AmericaBahia,
	tzlocation.AmericaBahiaBanderas,
	tzlocation.AmericaBarbados,
	tzlocation.AmericaBelem,
	tzlocation.AmericaBelize,
	tzlocation.AmericaBlancSablon,
	tzlocation.AmericaBoaVista,
	tzlocation.AmericaBogota,
	tzlocation.AmericaBoise,
	tzlocation.AmericaCambridgeBay,
	tzlocation.AmericaCampoGrande,
	tzlocation.AmericaCancun,
	tzlocation.AmericaCaracas,
	tzlocation.AmericaCayenne,
	tzlocation.AmericaCayman,
	tzlocation.AmericaChicago,
	tzlocation.AmericaChihuahua,
	tzlocation.AmericaCostaRica,
	tzlocation.AmericaCreston,
	tzlocation.AmericaCuiaba,
	tzlocation.AmericaCuracao,
	tzlocation.AmericaDanmarkshavn,
	tzlocation.AmericaDawson,
	tzlocation.AmericaDawsonCreek,
	tzlocation.AmericaDenver,
	tzlocation.AmericaDetroit,
	tzlocation.AmericaDominica,
	tzlocation.AmericaEdmonton,
	tzlocation.AmericaEirunepe,
	tzlocation.AmericaElSalvador,
	tzlocation.AmericaFortaleza,
	tzlocation.AmericaFortNelson,
	tzlocation.AmericaGlaceBay,
	tzlocation.AmericaGooseBay,
	tzlocation.AmericaGrandTurk,
	tzlocation.AmericaGrenada,
	tzlocation.AmericaGuadeloupe,
	tzlocation.AmericaGuatemala,
	tzlocation.AmericaGuayaquil,
	tzlocation.AmericaGuyana,
	tzlocation.AmericaHalifax,
	tzlocation.AmericaHavana,
	tzlocation.AmericaHermosillo,
	tzlocation.AmericaIndianaIndianapolis,
	tzlocation.AmericaIndianaKnox,
	tzlocation.AmericaIndianaMarengo,
	tzlocation.AmericaIndianaPetersburg,
	tzlocation.AmericaIndianaTellCity,
	tzlocation.AmericaIndianaVevay,
	tzlocation.AmericaIndianaVincennes,
	tzlocation.AmericaIndianaWinamac,
	tzlocation.AmericaInuvik,
	tzlocation.AmericaIqaluit,
	tzlocation.AmericaJamaica,
	tzlocation.AmericaJuneau,
	tzlocation.AmericaKentuckyLouisville,
	tzlocation.AmericaKentuckyMonticello,
	tzlocation.AmericaKralendijk,
	tzlocation.AmericaLaPaz,
	tzlocation.AmericaLima,
	tzlocation.AmericaLosAngeles,
	tzlocation.AmericaLowerPrinces,
	tzlocation.AmericaMaceio,
	tzlocation.AmericaManagua,
	tzlocation.AmericaManaus,
	tzlocation.AmericaMarigot,
	tzlocation.AmericaMartinique,
	tzlocation.AmericaMatamoros,
	tzlocation.AmericaMazatlan,
	tzlocation.AmericaMenominee,
	tzlocation.AmericaMerida,
	tzlocation.AmericaMetlakatla,
	tzlocation.AmericaMexicoCity,
	tzlocation.AmericaMiquelon,
	tzlocation.AmericaMoncton,
	tzlocation.AmericaMonterrey,
	tzlocation.AmericaMontevideo,
	tzlocation.AmericaMontserrat,
	tzlocation.AmericaNassau,
	tzlocation.AmericaNewYork,
	tzlocation.AmericaNipigon,
	tzlocation.AmericaNome,
	tzlocation.AmericaNoronha,
	tzlocation.AmericaNorthDakotaBeulah,
	tzlocation.AmericaNorthDakotaCenter,
	tzlocation.AmericaNorthDakotaNewSalem,
	tzlocation.AmericaNuuk,
	tzlocation.AmericaOjinaga,
	tzlocation.AmericaPanama,
	tzlocation.AmericaPangnirtung,
	tzlocation.AmericaParamaribo,
	tzlocation.AmericaPhoenix,
	tzlocation.AmericaPortauPrince,
	tzlocation.AmericaPortofSpain,
	tzlocation.AmericaPortoVelho,
	tzlocation.AmericaPuertoRico,
	tzlocation.AmericaPuntaArenas,
	tzlocation.AmericaRainyRiver,
	tzlocation.AmericaRankinInlet,
	tzlocation.AmericaRecife,
	tzlocation.AmericaRegina,
	tzlocation.AmericaResolute,
	tzlocation.AmericaRioBranco,
	tzlocation.AmericaSantarem,
	tzlocation.AmericaSantiago,
	tzlocation.AmericaSantoDomingo,
	tzlocation.AmericaSaoPaulo,
	tzlocation.AmericaScoresbysund,
	tzlocation.AmericaSitka,
	tzlocation.AmericaStBarthelemy,
	tzlocation.AmericaStJohns,
	tzlocation.AmericaStKitts,
	tzlocation.AmericaStLucia,
	tzlocation.AmericaStThomas,
	tzlocation.AmericaStVincent,
	tzlocation.AmericaSwiftCurrent,
	tzlocation.AmericaTegucigalpa,
	tzlocation.AmericaThule,
	tzlocation.AmericaThunderBay,
	tzlocation.AmericaTijuana,
	tzlocation.AmericaToronto,
	tzlocation.AmericaTortola,
	tzlocation.AmericaVancouver,
	tzlocation.AmericaWhitehorse,
	tzlocation.AmericaWinnipeg,
	tzlocation.AmericaYakutat,
	tzlocation.AmericaYellowknife,
	tzlocation.AntarcticaCasey,
	tzlocation.AntarcticaDavis,
	tzlocation.AntarcticaDumontDUrville,
	tzlocation.AntarcticaMacquarie,
	tzlocation.AntarcticaMawson,
	tzlocation.AntarcticaMcMurdo,
	tzlocation.AntarcticaPalmer,
	tzlocation.AntarcticaRothera,
	tzlocation.AntarcticaSyowa,
	tzlocation.AntarcticaTroll,
	tzlocation.AntarcticaVostok,
	tzlocation.ArcticLongyearbyen,
	tzlocation.AsiaAden,
	tzlocation.AsiaAlmaty,
	tzlocation.AsiaAmman,
	tzlocation.AsiaAnadyr,
	tzlocation.AsiaAqtau,
	tzlocation.AsiaAqtobe,
	tzlocation.AsiaAshgabat,
	tzlocation.AsiaAtyrau,
	tzlocation.AsiaBaghdad,
	tzlocation.AsiaBahrain,
	tzlocation.AsiaBaku,
	tzlocation.AsiaBangkok,
	tzlocation.AsiaBarnaul,
	tzlocation.AsiaBeirut,
	tzlocation.AsiaBishkek,
	tzlocation.AsiaBrunei,
	tzlocation.AsiaChita,
	tzlocation.AsiaChoibalsan,
	tzlocation.AsiaColombo,
	tzlocation.AsiaDamascus,
	tzlocation.AsiaDhaka,
	tzlocation.AsiaDili,
	tzlocation.AsiaDubai,
	tzlocation.AsiaDushanbe,
	tzlocation.AsiaFamagusta,
	tzlocation.AsiaGaza,
	tzlocation.AsiaHebron,
	tzlocation.AsiaHoChiMinh,
	tzlocation.AsiaHongKong,
	tzlocation.AsiaHovd,
	tzlocation.AsiaIrkutsk,
	tzlocation.AsiaJakarta,
	tzlocation.AsiaJayapura,
	tzlocation.AsiaJerusalem,
	tzlocation.AsiaKabul,
	tzlocation.AsiaKamchatka,
	tzlocation.AsiaKarachi,
	tzlocation.AsiaKathmandu,
	tzlocation.AsiaKhandyga,
	tzlocation.AsiaKolkata,
	tzlocation.AsiaKrasnoyarsk,
	tzlocation.AsiaKualaLumpur,
	tzlocation.AsiaKuching,
	tzlocation.AsiaKuwait,
	tzlocation.AsiaMacau,
	tzlocation.AsiaMagadan,
	tzlocation.AsiaMakassar,
	tzlocation.AsiaManila,
	tzlocation.AsiaMuscat,
	tzlocation.AsiaNicosia,
	tzlocation.AsiaNovokuznetsk,
	tzlocation.AsiaNovosibirsk,
	tzlocation.AsiaOmsk,
	tzlocation.AsiaOral,
	tzlocation.AsiaPhnomPenh,
	tzlocation.AsiaPontianak,
	tzlocation.AsiaPyongyang,
	tzlocation.AsiaQatar,
	tzlocation.AsiaQostanay,
	tzlocation.AsiaQyzylorda,
	tzlocation.AsiaRiyadh,
	tzlocation.AsiaSakhalin,
	tzlocation.AsiaSamarkand,
	tzlocation.AsiaSeoul,
	tzlocation.AsiaShanghai,
	tzlocation.AsiaSingapore,
	tzlocation.AsiaSrednekolymsk,
	tzlocation.AsiaTaipei,
	tzlocation.AsiaTashkent,
	tzlocation.AsiaTbilisi,
	tzlocation.AsiaTehran,
	tzlocation.AsiaThimphu,
	tzlocation.AsiaTokyo,
	tzlocation.AsiaTomsk,
	tzlocation.AsiaUlaanbaatar,
	tzlocation.AsiaUrumqi,
	tzlocation.AsiaUstNera,
	tzlocation.AsiaVientiane,
	tzlocation.AsiaVladivostok,
	tzlocation.AsiaYakutsk,
	tzlocation.AsiaYangon,
	tzlocation.AsiaYekaterinburg,
	tzlocation.AsiaYerevan,
	tzlocation.AtlanticAzores,
	tzlocation.AtlanticBermuda,
	tzlocation.AtlanticCanary,
	tzlocation.AtlanticCapeVerde,
	tzlocation.AtlanticFaroe,
	tzlocation.AtlanticMadeira,
	tzlocation.AtlanticReykjavik,
	tzlocation.AtlanticSouthGeorgia,
	tzlocation.AtlanticStanley,
	tzlocation.AtlanticStHelena,
	tzlocation.AustraliaAdelaide,
	tzlocation.AustraliaBrisbane,
	tzlocation.AustraliaBrokenHill,
	tzlocation.AustraliaCurrie,
	tzlocation.AustraliaDarwin,
	tzlocation.AustraliaEucla,
	tzlocation.AustraliaHobart,
	tzlocation.AustraliaLindeman,
	tzlocation.AustraliaLordHowe,
	tzlocation.AustraliaMelbourne,
	tzlocation.AustraliaPerth,
	tzlocation.AustraliaSydney,
	tzlocation.EuropeAmsterdam,
	tzlocation.EuropeAndorra,
	tzlocation.EuropeAstrakhan,
	tzlocation.EuropeAthens,
	tzlocation.EuropeBelgrade,
	tzlocation.EuropeBerlin,
	tzlocation.EuropeBratislava,
	tzlocation.EuropeBrussels,
	tzlocation.EuropeBucharest,
	tzlocation.EuropeBudapest,
	tzlocation.EuropeBusingen,
	tzlocation.EuropeChisinau,
	tzlocation.EuropeCopenhagen,
	tzlocation.EuropeDublin,
	tzlocation.EuropeGibraltar,
	tzlocation.EuropeGuernsey,
	tzlocation.EuropeHelsinki,
	tzlocation.EuropeIsleofMan,
	tzlocation.EuropeIstanbul,
	tzlocation.EuropeJersey,
	tzlocation.EuropeKaliningrad,
	tzlocation.EuropeKiev,
	tzlocation.EuropeKirov,
	tzlocation.EuropeLisbon,
	tzlocation.EuropeLjubljana,
	tzlocation.EuropeLondon,
	tzlocation.EuropeLuxembourg,
	tzlocation.EuropeMadrid,
	tzlocation.EuropeMalta,
	tzlocation.EuropeMariehamn,
	tzlocation.EuropeMinsk,
	tzlocation.EuropeMonaco,
	tzlocation.EuropeMoscow,
	tzlocation.EuropeOslo,
	tzlocation.EuropeParis,
	tzlocation.EuropePodgorica,
	tzlocation.EuropePrague,
	tzlocation.EuropeRiga,
	tzlocation.EuropeRome,
	tzlocation.EuropeSamara,
	tzlocation.EuropeSanMarino,
	tzlocation.EuropeSarajevo,
	tzlocation.EuropeSaratov,
	tzlocation.EuropeSimferopol,
	tzlocation.EuropeSkopje,
	tzlocation.EuropeSofia,
	tzlocation.EuropeStockholm,
	tzlocation.EuropeTallinn,
	tzlocation.EuropeTirane,
	tzlocation.EuropeUlyanovsk,
	tzlocation.EuropeUzhgorod,
	tzlocation.EuropeVaduz,
	tzlocation.EuropeVatican,
	tzlocation.EuropeVienna,
	tzlocation.EuropeVilnius,
	tzlocation.EuropeVolgograd,
	tzlocation.EuropeWarsaw,
	tzlocation.EuropeZagreb,
	tzlocation.EuropeZaporozhye,
	tzlocation.EuropeZurich,
	tzlocation.IndianAntananarivo,
	tzlocation.IndianChagos,
	tzlocation.IndianChristmas,
	tzlocation.IndianCocos,
	tzlocation.IndianComoro,
	tzlocation.IndianKerguelen,
	tzlocation.IndianMahe,
	tzlocation.IndianMaldives,
	tzlocation.IndianMauritius,
	tzlocation.IndianMayotte,
	tzlocation.IndianReunion,
	tzlocation.PacificApia,
	tzlocation.PacificAuckland,
	tzlocation.PacificBougainville,
	tzlocation.PacificChatham,
	tzlocation.PacificChuuk,
	tzlocation.PacificEaster,
	tzlocation.PacificEfate,
	tzlocation.PacificEnderbury,
	tzlocation.PacificFakaofo,
	tzlocation.PacificFiji,
	tzlocation.PacificFunafuti,
	tzlocation.PacificGalapagos,
	tzlocation.PacificGambier,
	tzlocation.PacificGuadalcanal,
	tzlocation.PacificGuam,
	tzlocation.PacificHonolulu,
	tzlocation.PacificKiritimati,
	tzlocation.PacificKosrae,
	tzlocation.PacificKwajalein,
	tzlocation.PacificMajuro,
	tzlocation.PacificMarquesas,
	tzlocation.PacificMidway,
	tzlocation.PacificNauru,
	tzlocation.PacificNiue,
	tzlocation.PacificNorfolk,
	tzlocation.PacificNoumea,
	tzlocation.PacificPagoPago,
	tzlocation.PacificPalau,
	tzlocation.PacificPitcairn,
	tzlocation.PacificPohnpei,
	tzlocation.PacificPortMoresby,
	tzlocation.PacificRarotonga,
	tzlocation.PacificSaipan,
	tzlocation.PacificTahiti,
	tzlocation.PacificTarawa,
	tzlocation.PacificTongatapu,
	tzlocation.PacificWake,
	tzlocation.PacificWallis,
}

func TestTZLocation_Load(t *testing.T) {
	for _, v := range data {
		_, err := v.Load()
		if err != nil {
			t.Errorf("tzlocation.TZLocation.Load wrong. location=%s err=%s", v, err.Error())
		}
	}
}
