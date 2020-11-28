import L from 'leaflet'

const iconPerson = new L.Icon({
	// after delete hard-code url need delete image using link https://postimg.cc/delete/L7RpR42d/1ebed331
	iconUrl: 'https://i.postimg.cc/TYMXZ8qJ/image.png',
	iconAnchor: undefined,
	popupAnchor: [2, -20],
	shadowUrl: undefined,
	shadowSize: undefined,
	shadowAnchor: undefined,
	iconSize: [50, 50],
	className: 'leaflet-div-icon'
});

export { iconPerson };
