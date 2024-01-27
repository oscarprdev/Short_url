import { useGlowCards } from '../hooks/useGlowCards';
import { Url } from '../types/url';
import UrlCard from './UrlCard';

const CONFIG = {
	proximity: 40,
	spread: 100,
	blur: 80,
	opacity: 0,
};

interface UrlCardsProps {
	urls: Url[];
	isHome?: boolean;
}

const UrlCardList = ({ urls, isHome }: UrlCardsProps) => {
	const { containerRef, cardsRef } = useGlowCards(CONFIG);

	return (
		<ul
			ref={containerRef}
			className='animate-fade-up link-container gap-5'>
			{urls.map((url) => (
				<UrlCard
					url={url}
					cardsRef={cardsRef}
					isHome={isHome}
				/>
			))}
		</ul>
	);
};

export default UrlCardList;
