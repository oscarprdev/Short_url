import { useErrorToast } from '../../hooks/useErrorToast';
import { useShortUrl } from '../../hooks/useShortUrl';
import { useGlobalStore } from '../../store/globalState';
import ShortUrlForm from '../ShortUrlForm';
import UrlCardList from '../urlCard/UrlCardList';

const UserScreenContainer = () => {
	const { user, urls, error } = useGlobalStore();
	const { mutate: addUrl } = useShortUrl();

	useErrorToast(error);

	return (
		<main data-testid="user-screen" className="flex flex-col gap-5 h-full sm:w-full items-center z-10">
			<div className="flex flex-col gap-5 items-center">
				<h2 className="text-[1.5rem] lg:text-[2rem]">Short your favourite links!</h2>
				<ShortUrlForm addUrl={addUrl} userId={user?.id} />
			</div>
			<p aria-label="link-counter" className="md:left-20 text-sm text-stone-400">
				Total {urls?.length}
			</p>
			{urls && <UrlCardList urls={urls} />}
		</main>
	);
};

export default UserScreenContainer;
