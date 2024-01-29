import { Url } from '@/src/types/url';
import { create } from 'zustand';
import { persist, createJSONStorage } from 'zustand/middleware';
import { User } from '../types/user';

interface storeState {
	user: User | null;
	urls: Url[] | null;
	error: string | null;
	setUser: (user: User) => void;
	setUrls: (urls: Url[]) => void;
	addUrl: (url: Url) => void;
	clearStore: () => void;
	setError: (error: string | null) => void;
}

export const useGlobalStore = create<storeState>()(
	persist(
		(set, get) => ({
			user: null,
			urls: null,
			error: null,
			setUser: (user) => set({ user }),
			setUrls: (urls) => set({ urls }),
			addUrl: (url) => set({ urls: [url, ...(get().urls || [])] }),
			clearStore: () => set({ urls: null, user: null, error: null }),
			setError: (error) => set({ error }),
		}),
		{
			name: 'user-store',
			storage: createJSONStorage(() => localStorage),
		}
	)
);
