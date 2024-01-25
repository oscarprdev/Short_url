import { useGlowCards } from '../hooks/useGlowCards';

import { Url } from '../types/url';
import UrlCardCopyIcon from './UrlCardCopyIcon';

const CONFIG = {
	proximity: 40,
	spread: 80,
	blur: 80,
	gap: 32,
	opacity: 0,
};

interface HomeUrlCardProps {
	url: Url;
}

const HomeUrlCard = ({ url }: HomeUrlCardProps) => {
	const { containerRef, cardsRef } = useGlowCards(CONFIG);

	return (
		<div
			ref={containerRef}
			className='animate-fade-up link-container '>
			<article
				ref={(el: HTMLDivElement | null) => el && cardsRef.current.push(el!)}
				className='relative link-card w-fit h-fit border border-stone-800'>
				<UrlCardCopyIcon url={url.shortUrl} />
				<a
					href={url.originalUrl}
					target='blank'
					className='font-light text-stone-300 hover:underline'>
					{url.shortUrl}
				</a>
				<div className='glows'></div>
			</article>
		</div>
	);
};

export default HomeUrlCard;
