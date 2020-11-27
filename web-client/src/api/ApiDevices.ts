import IDevice from '../interfaces/IDevice'

export const ApiSendDevice = async (device: IDevice) => {
	try {
		const res = await fetch('/api/devices', {
			method: 'POST',
			headers: {'Content-Type': 'application/json'},
			body: JSON.stringify(device)
		});
	} catch (err) {
		throw err;
	}
}
