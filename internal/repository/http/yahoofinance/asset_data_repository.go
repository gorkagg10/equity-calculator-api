package yahoofinance

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorkagg10/equity-calculator-api/internal/domain"
)

type AssetDataRepository struct {
}

func NewAssetDataRepository() AssetDataRepository {
	return AssetDataRepository{}
}

func (r AssetDataRepository) GetAssetData(symbol string) (*domain.AssetData, error) {
	// Usamos el endpoint de gráficos pidiendo solo 1 día de rango y 1 minuto de intervalo
	url := fmt.Sprintf("https://query1.finance.yahoo.com/v8/finance/chart/%s?range=1d&interval=1m", symbol)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creando petición: %w", err)
	}

	// CRÍTICO: Yahoo bloquea las peticiones que no tienen un User-Agent de navegador
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "application/json")

	// Cliente con timeout para evitar que la Goroutine se cuelgue
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error de red: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("yahoo devolvió código de error HTTP: %d", resp.StatusCode)
	}

	var yahooData YahooChartResponse
	if err := json.NewDecoder(resp.Body).Decode(&yahooData); err != nil {
		return nil, fmt.Errorf("decoding JSON: %w", err)
	}

	if yahooData.Chart.Error != nil {
		return nil, fmt.Errorf("yahoo error: %s - %s", yahooData.Chart.Error.Code, yahooData.Chart.Error.Description)
	}

	// Validar que hay resultados
	if len(yahooData.Chart.Result) == 0 {
		return nil, fmt.Errorf("no data for this symbol: %s", symbol)
	}

	// Extraer los metadatos
	meta := yahooData.Chart.Result[0].Meta
	// timestamp := time.Unix(meta.RegularMarketTime, 0).Format("2006-01-02 15:04:05")

	return domain.NewAssetData(
		meta.Symbol,
		meta.Name,
		meta.RegularMarketPrice,
		meta.Currency,
		meta.ExchangeName,
	), nil
}
