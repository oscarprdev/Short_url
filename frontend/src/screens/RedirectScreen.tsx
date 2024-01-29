import { useUrlUsage } from '../hooks/useUrlUsage';
import { useGlobalStore } from '../store/globalState';

interface RedirectScreenProps {
	urlId: string;
}

const RedirectScreen = ({ urlId }: RedirectScreenProps) => {
	const { urls } = useGlobalStore();
	const { mutate: useUrl } = useUrlUsage();

	useUrl({ urlId });

	const originalUrl = urls?.find((url) => url.id === urlId);

	window.location.href = originalUrl?.originalUrl || '/';
};

export default RedirectScreen;
