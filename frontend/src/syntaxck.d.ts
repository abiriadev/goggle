declare global {
	export interface Window {
		Go: any
		syntaxck: (query: string) => bool
	}
}

export {}
