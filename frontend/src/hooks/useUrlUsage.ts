import { urlUsage } from '../services/api/urlUsage';
import { useGlobalStore } from '../store/globalState';
import { useMutation } from '@tanstack/react-query';

export const useUrlUsage = () => {
	const { user, urls, setUrls, setError } = useGlobalStore();

	return useMutation({
		mutationFn: async ({ urlId }: { urlId: string }) => {
			return await urlUsage({ urlId, userId: user?.id });
		},
		onSuccess: ({ url }) => {
			if (urls) {
				const updatedUrls = urls.map(currentUrl => (currentUrl.id === url.id ? url : currentUrl));
				setUrls(updatedUrls);
			} else {
				setError('No urls availables');
			}
		},
		onError: error => {
			setError(error.message);
		},
	});
};
