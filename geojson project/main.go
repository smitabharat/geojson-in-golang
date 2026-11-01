package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"os"

// 	"github.com/gofiber/fiber/v2"
// 	geojson "github.com/paulmach/go.geojson"
// )

// func main() {
// 	url := "https://raw.githubusercontent.com/india-in-data/india-states-2019/master/india_states.geojson"

// 	resp, err := http.Get(url)
// 	if err != nil {
// 		fmt.Printf("Error fetching GeoJSON: %v\n", err)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		fmt.Printf("Non-OK HTTP status: %d\n", resp.StatusCode)
// 		body, _ := io.ReadAll(resp.Body)
// 		fmt.Printf("Response body: %s\n", string(body))
// 		return
// 	}

// 	data, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Printf("Error reading body: %v\n", err)
// 		return
// 	}

// 	// Unmarshal into a FeatureCollection
// 	fc, err := geojson.UnmarshalFeatureCollection(data)
// 	if err != nil {
// 		fmt.Printf("Error parsing GeoJSON: %v\n", err)
// 		return
// 	}

// 	fmt.Printf("Loaded FeatureCollection with %d features\n", len(fc.Features))

// 	outFile := "india_states_loaded.geojson"
// 	if err := os.WriteFile(outFile, data, 0644); err != nil {
// 		fmt.Printf("Error writing file: %v\n", err)
// 		return
// 	}

//		fmt.Println("GeoJSON saved to", outFile)
//	}
// func main() {
// 	app := fiber.New()

// 	// Endpoint to serve GeoJSON
// 	app.Get("/india_states_loaded.geojson", func(c *fiber.Ctx) error {
// 		fileName := "india_states_loaded.geojson"

// 		// If file not found locally, fetch from GitHub
// 		if _, err := os.Stat(fileName); os.IsNotExist(err) {
// 			url := "https://raw.githubusercontent.com/india-in-data/india-states-2019/master/india_states_loaded.geojson"
// 			resp, err := http.Get(url)
// 			if err != nil {
// 				return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Error fetching GeoJSON: %v", err))
// 			}
// 			defer resp.Body.Close()

// 			if resp.StatusCode != http.StatusOK {
// 				body, _ := io.ReadAll(resp.Body)
// 				return c.Status(resp.StatusCode).SendString(string(body))
// 			}

// 			data, err := io.ReadAll(resp.Body)
// 			if err != nil {
// 				return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Error reading GeoJSON: %v", err))
// 			}

// 			// Save to file for reuse
// 			os.WriteFile(fileName, data, 0644)
// 		}

// 		// Read local GeoJSON file
// 		data, err := os.ReadFile(fileName)
// 		if err != nil {
// 			return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Error reading local GeoJSON: %v", err))
// 		}

// 		// Unmarshal into FeatureCollection
// 		fc, err := geojson.UnmarshalFeatureCollection(data)
// 		if err != nil {
// 			return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Error parsing GeoJSON: %v", err))
// 		}

// 		// Convert to pretty JSON
// 		out, _ := json.MarshalIndent(fc, "", "  ")
// 		return c.Type("application/json").Send(out)
// 	})

// 	// Root route
// 	app.Get("/", func(c *fiber.Ctx) error {
// 		return c.SendString("🌏 Go Fiber India GeoJSON API Running! Visit /india-geojson")
// 	})

// 	// Start server
// 	fmt.Println("✅ Server running on http://localhost:3000")
// 	app.Listen(":3000")
// }
//package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	geojson "github.com/paulmach/go.geojson"
)

type City struct {
	Name   string    `json:"name"`
	Coords []float64 `json:"coords"`
	Icon   string    `json:"icon"`
}

func toJSON(v interface{}) string {
	b, _ := json.Marshal(v)
	return string(b)
}

func main() {
	app := fiber.New()

	// Serve GeoJSON
	app.Get("/india_states.geojson", func(c *fiber.Ctx) error {
		fileName := "india_states.geojson"
		url := "https://raw.githubusercontent.com/india-in-data/india-states-2019/master/india_states.geojson"

		if _, err := os.Stat(fileName); os.IsNotExist(err) {
			resp, err := http.Get(url)
			if err != nil {
				return c.Status(500).SendString(fmt.Sprintf("Error fetching GeoJSON: %v", err))
			}
			defer resp.Body.Close()

			data, _ := io.ReadAll(resp.Body)
			os.WriteFile(fileName, data, 0644)
		}

		data, err := os.ReadFile(fileName)
		if err != nil {
			return c.Status(500).SendString(fmt.Sprintf("Error reading file: %v", err))
		}

		var fc geojson.FeatureCollection
		if err := json.Unmarshal(data, &fc); err != nil {
			return c.Status(500).SendString(fmt.Sprintf("Error parsing GeoJSON: %v", err))
		}

		out, _ := json.MarshalIndent(fc, "", "  ")
		return c.Type("application/json").Send(out)
	})

	// API for cities list
	app.Get("/api/cities", func(c *fiber.Ctx) error {
		cities := []map[string]interface{}{
			{"name": "Delhi", "coords": []float64{28.6139, 77.2090}, "icon": "⭐"},
			{"name": "Mumbai", "coords": []float64{19.0760, 72.8777}, "icon": "🏙️"},
			{"name": "Kolkata", "coords": []float64{22.5726, 88.3639}, "icon": "📍"},
			{"name": "Chennai", "coords": []float64{13.0827, 80.2707}, "icon": "🌆"},
			{"name": "Bengaluru", "coords": []float64{12.9716, 77.5946}, "icon": "🚀"},
			{"name": "Hyderabad", "coords": []float64{17.3850, 78.4867}, "icon": "🌟"},
			{"name": "Pune", "coords": []float64{18.5204, 73.8567}, "icon": "📌"},
			{"name": "Jaipur", "coords": []float64{26.9124, 75.7873}, "icon": "🏰"},
		}
		return c.JSON(cities)
	})

	// Serve map with dropdown
	app.Get("/", func(c *fiber.Ctx) error {
		html := `
<!DOCTYPE html>
<html>
<head>
<title>India Cities Map</title>
<meta charset="utf-8" />
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<link rel="stylesheet" href="https://unpkg.com/leaflet@1.9.4/dist/leaflet.css" />
<script src="https://unpkg.com/leaflet@1.9.4/dist/leaflet.js"></script>
<style>
  body { margin: 0; }
  #map { height: 100vh; width: 100%; }
  #dropdown-container {
    position: absolute;
    top: 10px; left: 10px;
    z-index: 1000;
    background: white;
    padding: 10px;
    border-radius: 8px;
    box-shadow: 0 2px 5px rgba(0,0,0,0.3);
  }
</style>
</head>
<body>
<div id="dropdown-container">
  <label for="citySelect">Select City:</label>
  <select id="citySelect">
    <option value="">-- All Cities --</option>
  </select>
</div>
<div id="map"></div>

<script>
  var map = L.map('map').setView([22.9734, 78.6569], 5);

  L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
    maxZoom: 18,
    attribution: '© OpenStreetMap contributors'
  }).addTo(map);

  // Load India GeoJSON
  fetch('/india_states.geojson')
    .then(res => res.json())
    .then(data => {
      L.geoJSON(data, {
        style: { color: '#3388ff', weight: 1, fillOpacity: 0.1 }
      }).addTo(map);
    });

  var cityMarkers = [];

  // Load cities from API
  fetch('/api/cities')
    .then(res => res.json())
    .then(cities => {
      var select = document.getElementById("citySelect");

      cities.forEach(city => {
        // Add option to dropdown
        var opt = document.createElement("option");
        opt.value = city.name;
        opt.textContent = city.name;
        select.appendChild(opt);

        // Add marker for each city with unique icon
        var marker = L.marker(city.coords, {
          title: city.name,
          icon: L.divIcon({
            className: "city-icon",
            html: "<div style='font-size:22px;'>"+city.icon+"</div>",
            iconSize: [25, 25],
            iconAnchor: [12, 12]
          })
        }).addTo(map)
        .bindPopup("<b>" + city.name + "</b><br>Latitude: " + city.coords[0] + "<br>Longitude: " + city.coords[1]);

        cityMarkers.push({ name: city.name, marker: marker, coords: city.coords });
      });

      // When city is selected
      select.addEventListener("change", function() {
        var selected = this.value;
        if (selected) {
          var city = cityMarkers.find(c => c.name === selected);
          if (city) {
            map.setView(city.coords, 8);
            city.marker.openPopup();
          }
        } else {
          map.setView([22.9734, 78.6569], 5); // Reset to India view
        }
      });
    });
</script>
</body>
</html>`

		return c.Type("html").SendString(html)
	})

	fmt.Println("✅ Server running at http://localhost:3000")
	fmt.Println("👉 Try opening in browser: http://localhost:3000")
	app.Listen(":3000")
}
