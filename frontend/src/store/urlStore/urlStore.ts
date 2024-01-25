import { create } from 'zustand';

interface UrlStore {
	url: string;
	setUrl: (url: string) => void;
}

export const useUrlStore = create<UrlStore>((set) => ({
	url: '',
	setUrl: (newUrl: string) => set({ url: newUrl }),
}));
