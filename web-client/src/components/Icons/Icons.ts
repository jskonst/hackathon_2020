import L from 'leaflet'

const iconPerson = new L.Icon({
	// this image will be deleted 5 or 6 December
	iconUrl: 'https://i.postimg.cc/brs5Gfkt/cat.jpg',
	iconAnchor: undefined,
	popupAnchor: [2, -20],
	iconSize: [50, 50],
	className: 'leaflet-div-icon'
});

export { iconPerson };
