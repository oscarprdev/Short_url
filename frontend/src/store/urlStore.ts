import { Url } from '@/src/types/url';
import { create } from 'zustand';
import { persist, createJSONStorage } from 'zustand/middleware';

interface UrlStoreState {
	urls: Url[] | null;
	setUrl: (url: Url) => void;
}

export const useUrlStore = create<UrlStoreState>()(
	persist(
		(set) => ({
			urls: null,
			setUrl: (url) => set({ urls: [url] }),
		}),
		{
			name: 'url-store',
			storage: createJSONStorage(() => localStorage),
		}
	)
);
