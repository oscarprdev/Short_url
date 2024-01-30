import { useQuery } from '@tanstack/react-query';
import { getUserInfo } from '../services/api/getUserInfo';
import { useGlobalStore } from '../store/globalState';
import { useEffect } from 'react';
import { API_URL } from '../constants/apiUrl';

export const useUserLogged = (userId: string) => {
	const { setUrls, setUser, clearStore, setError } = useGlobalStore();

	const { data, isLoading, isError } = useQuery({
		queryKey: ['user'],
		queryFn: async () => await getUserInfo({ userId }),
		staleTime: Infinity,
	});

	useEffect(() => {
		if (!isLoading && data?.status === 200) {
			setUrls(data.urls);
			setUser(data.user);
		}

		if (data?.status && data?.status !== 200) {
			window.location.href = `${API_URL}/auth/logout`;
			clearStore();
			setError('User unauthorized or token expired');
		}
		// eslint-disable-next-line react-hooks/exhaustive-deps
	}, [data, isLoading, setUrls, setUser]);

	return { isLoading, isError };
};
