package main

import (
    "fmt"
)

func main() {
    const maandenPerJaar = 12

    var maandelijksInkomen float64 = 2140.0
    var extraJaarlijksInkomen float64 = 4000.0
    var vasteMaandelijkseKost float64 = 400.0
    var geplandeJaarlijkseUitgaven float64 = 10000
    var sparenPerJaar float64 = 10000

    vereistJaarlijksSparen := sparenPerJaar - extraJaarlijksInkomen
    if vereistJaarlijksSparen < 0 {
        geplandeJaarlijkseUitgaven += -vereistJaarlijksSparen
        vereistJaarlijksSparen = 0
    }

    totaalJaarlijksInkomen := maandelijksInkomen * maandenPerJaar
    totaleUitgaven := vasteMaandelijkseKost * maandenPerJaar + geplandeJaarlijkseUitgaven
    maandelijksVrijTeBesteden := (totaalJaarlijksInkomen - totaleUitgaven - vereistJaarlijksSparen) / maandenPerJaar

    fmt.Printf("\nMaandelijks naar Vrij Te Besteden Rekening: €%.2f\n", maandelijksVrijTeBesteden)
    fmt.Printf("Maandelijks laten staan op Zichtrekening (Vaste verplichte kosten): €%.2f\n", vasteMaandelijkseKost)
    fmt.Printf("Maandelijks naar Geplande Spaarrekening: €%.2f\n", geplandeJaarlijkseUitgaven/maandenPerJaar)
    fmt.Printf("Maandelijks naar Lange Termijn Spaarrekening: €%.2f\n", vereistJaarlijksSparen/maandenPerJaar)
    fmt.Printf("Jaarlijkse bedrag extra sparen op Lange Termijn Spaarrekening: €%.2f\n", vereistJaarlijksSparen)
}
