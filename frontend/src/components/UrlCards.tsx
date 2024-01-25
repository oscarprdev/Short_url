import { useGlowCards } from '../hooks/useGlowCards';

import { Url } from '../types/url';
import { formatTimeStamp } from '../utils/formatTimeStamp';
import UrlCardCopyIcon from './UrlCardCopyIcon';

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

const UrlCards = ({ urls, isHome }: UrlCardsProps) => {
	const { containerRef, cardsRef } = useGlowCards(CONFIG);

	return (
		<div
			ref={containerRef}
			className='animate-fade-up link-container '>
			{urls.map((url) => (
				<article
					key={url.id}
					ref={(el: HTMLDivElement | null) => el && cardsRef.current.push(el!)}
					className='relative flex flex-col gap-4 link-card w-fit h-fit border border-stone-800 font-light text-stone-300'>
					<a
						href={url.originalUrl}
						target='blank'
						className='hover:underline'>
						{url.shortUrl}
					</a>
					{!isHome && (
						<>
							<UrlCardCopyIcon url={url.shortUrl} />
							<div className='flex flex-col text-sm text-stone-400'>
								<p>Times used: {url.usage}</p>
								<p>Expiration: {formatTimeStamp(url.expiresAt)}</p>
							</div>
						</>
					)}
					<div className='glows'></div>
				</article>
			))}
		</div>
	);
};

export default UrlCards;
