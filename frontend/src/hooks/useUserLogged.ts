import { useQuery } from '@tanstack/react-query';
import { getUserInfo } from '../services/api/getUserInfo';
import { useUserStore } from '../store/userStore';
import { useEffect } from 'react';

export const useUserLogged = (userId: string) => {
	const { setUrls, setUser } = useUserStore();

	const { data, isLoading, isError } = useQuery({
		queryKey: ['user'],
		queryFn: async () => await getUserInfo({ userId }),
	});

	useEffect(() => {
		if (!isLoading && data?.status === 200) {
			setUrls(data.urls);
			setUser(data.user);
		}
	}, [data, isLoading, setUrls, setUser]);

	if (data?.status !== 200) {
		return { isLoading, isError: true };
	}

	return { isLoading, isError };
};
