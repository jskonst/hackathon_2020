import L from 'leaflet'

const iconPerson = new L.Icon({
	iconUrl: require('../img/image.png'),
	iconRetinaUrl: require('../img/image.png'),
	iconAnchor: undefined,
	popupAnchor: undefined, // new L.Point(50, 50),
	shadowUrl: undefined,
	shadowSize: undefined,
	shadowAnchor: undefined,
	iconSize: new L.Point(60, 60),
	className: 'leaflet-div-icon'
});

export { iconPerson };
