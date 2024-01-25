import { Url } from '@/src/types/url';
import { create } from 'zustand';

interface UrlStore {
	url: Url | null;
	setUrl: (input: Url) => void;
}

export const useUrlStore = create<UrlStore>((set) => ({
	url: null,
	setUrl: (url) => set({ url }),
}));
