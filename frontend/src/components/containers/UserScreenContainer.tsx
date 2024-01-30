import ShortUrlForm from '../ShortUrlForm';
import UrlCardList from '../urlCard/UrlCardList';
import { useShortUrl } from '../../hooks/useShortUrl';
import { useGlobalStore } from '../../store/globalState';
import { useState } from 'react';
import CardListNav from '../CardListNav';
import { useErrorToast } from '../../hooks/useErrorToast';

const UserScreenContainer = () => {
	const [isRowsLayout, setIsRowsLayout] = useState(false);

	const { user, urls, error } = useGlobalStore();
	const { mutate: addUrl } = useShortUrl();

	useErrorToast(error);

	const onSetRowLayout = () => {
		setIsRowsLayout(true);
	};

	const onSetGridLayout = () => {
		setIsRowsLayout(false);
	};

	return (
		<main
			data-testid='user-screen'
			className='flex flex-col gap-5 h-full sm:w-full items-center z-10'>
			<div className='flex flex-col gap-5 items-center'>
				<h2
					aria-label='call to action text'
					className='text-[1.5rem] lg:text-[2rem]'>
					Short your favourite links!
				</h2>
				<ShortUrlForm
					addUrl={addUrl}
					userId={user?.id}
				/>
			</div>
			<CardListNav
				isRowsLayout={isRowsLayout}
				onSetRowLayout={onSetRowLayout}
				onSetGridLayout={onSetGridLayout}
			/>
			{urls && (
				<UrlCardList
					urls={urls}
					isRowsLayout={isRowsLayout}
				/>
			)}
		</main>
	);
};

export default UserScreenContainer;
