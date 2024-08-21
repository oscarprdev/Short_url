import { Url } from '../../types/url';
import UrlCard from './UrlCard';
import UserUrlCard from './UserUrlCard';

interface UrlCardsProps {
	urls: Url[];
	isHome?: boolean;
}

const UrlCardList = ({ urls, isHome }: UrlCardsProps) => {
	return (
		<ul
			className={`${'w-[90vw] lg:w-[1100px]'} animate-fade-up pt-2 pb-20 link-container flex-wrap items-start gap-5 h-fit`}>
			{urls.map(url => (
				<>{isHome ? <UrlCard key={url.id} url={url} /> : <UserUrlCard key={url.id} url={url} />}</>
			))}
		</ul>
	);
};

export default UrlCardList;
