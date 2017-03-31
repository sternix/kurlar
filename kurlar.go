package kurlar

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type (
	Kur struct {
		XMLName    xml.Name   `xml:"Tarih_Date"`
		Tarih      string     `xml:"Tarih,attr"`
		Date       string     `xml:"Date,attr"`
		BultenNo   string     `xml:"Bulten_No,attr"`
		Currencies []Currency `xml:"Currency"`
	}

	Currency struct {
		CrossOrder      string `xml:"CrossOrder,attr"`
		Kod             string `xml:"Kod,attr"`
		CurrencyCode    string `xml:"CurrencyCode,attr"`
		Unit            string `xml:"Unit"`
		Isim            string `xml:"Isim"`
		CurrencyName    string `xml:"CurrencyName"`
		ForexBuying     string `xml:"ForexBuying"`
		ForexSelling    string `xml:"ForexSelling"`
		BanknoteBuying  string `xml:"BanknoteBuying"`
		BanknoteSelling string `xml:"BanknoteSelling"`
		CrossRateUSD    string `xml:"CrossRateUSD"`
		CrossRateOther  string `xml:"CrossRateOther"`
	}
)

// değiştirilebilir
var TcmbKurlarXmlUrl = "http://www.tcmb.gov.tr/kurlar/today.xml"

func Today() (*Kur, error) {
	file, err := fetchKurlarXml()
	if err != nil {
		return nil, err
	}

	defer file.Close()
	var kur Kur
	if err := xml.NewDecoder(file).Decode(&kur); err != nil {
		return nil, fmt.Errorf("Xml dosyasını okurken hata oluştu: %s", err)
	}

	return &kur, nil
}

func fetchKurlarXml() (*os.File, error) {
	resp, err := http.Get(TcmbKurlarXmlUrl)
	if err != nil {
		return nil, fmt.Errorf("kurlar.xml alınamadı: %s", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("kurlar.xml alınamadı: StatusCode = %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Alınan dosya okunamadı: %s", err)
	}

	file, err := ioutil.TempFile("", "today.xml")
	if err != nil {
		return nil, fmt.Errorf("Temp dosya oluşturulamadı: %s", err)
	}

	file.Write(body)
	file.Seek(0, 0)
	return file, nil
}

func (k *Kur) String() string {
	buf := &bytes.Buffer{}

	fmt.Fprintf(buf, "Tarih: %s\n", k.Tarih)
	fmt.Fprintf(buf, "Date: %s\n", k.Date)
	fmt.Fprintf(buf, "BultenNo: %s\n", k.BultenNo)
	for _, c := range k.Currencies {
		fmt.Fprintln(buf)
		fmt.Fprintf(buf, "\tCrossOrder: %s\n", c.CrossOrder)
		fmt.Fprintf(buf, "\tKod: %s\n", c.Kod)
		fmt.Fprintf(buf, "\tCurrencyCode: %s\n", c.CurrencyCode)
		fmt.Fprintf(buf, "\tUnit: %s\n", c.Unit)
		fmt.Fprintf(buf, "\tIsim: %s\n", c.Isim)
		fmt.Fprintf(buf, "\tCurrencyName: %s\n", c.CurrencyName)
		fmt.Fprintf(buf, "\tForexBuying: %s\n", c.ForexBuying)
		fmt.Fprintf(buf, "\tForexSelling: %s\n", c.ForexSelling)
		fmt.Fprintf(buf, "\tBanknoteBuying: %s\n", c.BanknoteBuying)
		fmt.Fprintf(buf, "\tBanknoteSelling: %s\n", c.BanknoteSelling)
		fmt.Fprintf(buf, "\tCrossRateUSD: %s\n", c.CrossRateUSD)
		fmt.Fprintf(buf, "\tCrossRateOther: %s\n", c.CrossRateOther)
	}

	return buf.String()
}
