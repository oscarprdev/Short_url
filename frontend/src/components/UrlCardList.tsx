import { useGlowCards } from '../hooks/useGlowCards';
import { Url } from '../types/url';
import UrlCard from './UrlCard';
import UserUrlCard from './UserUrlCard';

const CONFIG = {
	proximity: 40,
	spread: 100,
	blur: 80,
	opacity: 0,
};

interface UrlCardsProps {
	urls: Url[];
	isHome?: boolean;
	isRowsLayout?: boolean;
}

const UrlCardList = ({ urls, isHome, isRowsLayout }: UrlCardsProps) => {
	const { containerRef, cardsRef } = useGlowCards(CONFIG);

	return (
		<>
			<ul
				ref={containerRef}
				className={`${isRowsLayout ? 'w-[800px]' : 'w-[1100px]'} animate-fade-up link-container flex-wrap items-start gap-5  mt-[-1.5rem]`}>
				{urls.map((url) => {
					return isHome ? (
						<UrlCard
							url={url}
							cardsRef={cardsRef}
						/>
					) : (
						<UserUrlCard
							url={url}
							cardsRef={cardsRef}
							isRowsLayout={isRowsLayout}
						/>
					);
				})}
			</ul>
		</>
	);
};

export default UrlCardList;
