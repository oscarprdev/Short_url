import { Url } from '@/src/types/url';
import { create } from 'zustand';
import { persist, createJSONStorage } from 'zustand/middleware';
import { User } from '../types/user';

interface UserStoreState {
	user: User | null;
	urls: Url[] | null;
	setUser: (user: User) => void;
	setUrls: (urls: Url[]) => void;
	addUrl: (url: Url) => void;
}

export const useUserStore = create<UserStoreState>()(
	persist(
		(set, get) => ({
			user: null,
			urls: null,
			setUser: (user) => set({ user }),
			setUrls: (urls) => set({ urls }),
			addUrl: (url) => set({ urls: [url, ...(get().urls || [])] }),
		}),
		{
			name: 'user-store',
			storage: createJSONStorage(() => localStorage),
		}
	)
);
