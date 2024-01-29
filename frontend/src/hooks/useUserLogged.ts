import { useQuery } from '@tanstack/react-query';
import { getUserInfo } from '../services/api/getUserInfo';
import { useGlobalStore } from '../store/globalState';
import { useEffect } from 'react';

export const useUserLogged = (userId: string) => {
	const { setUrls, setUser, clearStore } = useGlobalStore();

	const { data, isLoading, isError } = useQuery({
		queryKey: ['user'],
		queryFn: async () => await getUserInfo({ userId }),
		staleTime: Infinity,
	});

	useEffect(() => {
		console.log(data);
		if (!isLoading && data?.status === 200) {
			setUrls(data.urls);
			setUser(data.user);
		}

		if (data?.status && data?.status !== 200) {
			window.location.href = 'http://localhost:8080/auth/logout';
			clearStore();
		}
	}, [data, isLoading, setUrls, setUser]);

	return { isLoading, isError };
};
