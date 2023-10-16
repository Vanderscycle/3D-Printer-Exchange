export interface Badge {
	msg: string;
	color: string;
	text: string;
}

export interface NavigationButton {
	url: string;
	name: string;
}

export interface User {
    username: string;
    email: string; // needs to be refined (e.g. make a valid email)
    password: string; // obciously we do not want their password

}

export interface Printer {
    user: string //User
    brand: string
    name: string
    powerSupply: string
    probe: string | null
    board: string
    hotend: string
    extruder: string
    nozzle: string
    buildVolume: string
    mods: string | object // needs to be an object (changed in backend required)
}
export interface Mods {
    name: string
    location: string
}
