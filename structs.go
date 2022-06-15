package flightaware

type Response struct {
	Flights []Flight `json:"flights"`
}

type Flight struct {
	Ident               string   `json:"ident"`
	IdentIcao           string   `json:"ident_icao"`
	IdentIata           string   `json:"ident_iata"`
	FaFlightID          string   `json:"fa_flight_id"`
	Operator            string   `json:"operator"`
	OperatorIcao        string   `json:"operator_icao"`
	OperatorIata        string   `json:"operator_iata"`
	FlightNumber        string   `json:"flight_number"`
	Registration        string   `json:"registration"`
	AtcIdent            string   `json:"atc_ident"`
	InboundFaFlightID   string   `json:"inbound_fa_flight_id"`
	Codeshares          []string `json:"codeshares"`
	CodesharesIata      []string `json:"codeshares_iata"`
	Blocked             bool     `json:"blocked"`
	Diverted            bool     `json:"diverted"`
	Cancelled           bool     `json:"cancelled"`
	PositionOnly        bool     `json:"position_only"`
	Origin              Airport  `json:"origin"`
	Destination         Airport  `json:"destination"`
	DepartureDelay      int      `json:"departure_delay"`
	ArrivalDelay        int      `json:"arrival_delay"`
	FiledETE            int      `json:"filed_ete"`
	ScheduledOut        string   `json:"scheduled_out"`
	EstimatedOut        string   `json:"estimated_out"`
	ActualOut           string   `json:"actual_out"`
	ScheduledOff        string   `json:"scheduled_off"`
	EstimatedOff        string   `json:"estimated_off"`
	ActualOff           string   `json:"actual_off"`
	ScheduledOn         string   `json:"scheduled_on"`
	EstimatedOn         string   `json:"estimated_on"`
	ActualOn            string   `json:"actual_on"`
	ScheduledIn         string   `json:"scheduled_in"`
	EstimatedIn         string   `json:"estimated_in"`
	ActualIn            string   `json:"actual_in"`
	Progress            int      `json:"progress_percent"`
	Status              string   `json:"status"`
	AircraftType        string   `json:"aircraft_type"`
	RouteDistance       int      `json:"route_distance"`
	FiledAirspeed       int      `json:"filed_airspeed"`
	FiledAltitude       int      `json:"filed_altitude"`
	Route               string   `json:"route"`
	BaggageClaim        string   `json:"baggage_claim"`
	BusinessClass       int      `json:"seats_cabin_business"`
	CoachClass          int      `json:"seats_cabin_coach"`
	FirstClass          int      `json:"seats_cabin_first"`
	GateOrigin          string   `json:"gate_origin"`
	GateDestination     string   `json:"gate_destination"`
	TerminalOrigin      string   `json:"terminal_origin"`
	TerminalDestination string   `json:"terminal_destination"`
	FlightType          string   `json:"type"`
}

type Airport struct {
	Code           string `json:"code"`
	CodeIcao       string `json:"code_icao"`
	CodeIata       string `json:"code_iata"`
	CodeLid        string `json:"code_lid"`
	AirportInfoUrl string `json:"airport_info_url"`
}
