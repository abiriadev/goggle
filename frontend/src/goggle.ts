export interface ResultSet {
	items: Array<ResultSet>
}

export interface ResultItem {
	sim: number
	sig: string
	summary: string
	link: string
}

const host = '/api'

export const query = async (
	query: string,
): Promise<ResultSet> => {
	const raw = await fetch(`${host}/query`, {
		method: 'POST',
		body: JSON.stringify({ q: query }),
	})

	return await raw.json()
}
