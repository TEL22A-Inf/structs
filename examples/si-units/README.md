# Si-Einheitensystem

Hier wird ein einfaches Einheitensystem mittels Structs definiert.

* [`Duration`](duration.go) beschreibt Zeiträume.
  Eine `Duration` kann auf verschiedene Weisen aus
  Sekunden-, Minuten-, Tages- etc. Angaben erzeugt werden.
* [`Distance`](distance.go) beschreibt Entfernungen.
  Eine `Distance` kann auf verschiedene Weisen aus
  Meter-, Kilometer-, Meilen- etc. Angaben erzeugt werden.
* [`Velocity`](velocity.go) beschreibt Geschwindigkeiten.
  Eine `Velocity` kann entweder aus `Duration`- und `Distance`-Objekten
  oder aus Zahlenangaben gebildet werden.
