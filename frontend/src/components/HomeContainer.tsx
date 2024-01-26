import { useShortUrl } from '../hooks/useShortUrl';
import { useUrlStore } from '../store/urlStore';
import ButtonLogin from './ButtonLogin';
import ShortUrlForm from './ShortUrlForm';
import UrlCards from './UrlCards';

const HomeContainer = () => {
	const { urls } = useUrlStore();
	const { mutate: addurl } = useShortUrl();

	return (
		<>
			<h1 className='text-[6rem] font-geistUltra'>Short - it</h1>
			<div className='flex flex-col items-center gap-3'>
				<h2 className='text-[1.2rem]'>The simplest URL Shortner you were waiting for.</h2>
				<ShortUrlForm addUrl={addurl} />
				{urls && (
					<UrlCards
						urls={urls}
						isHome
					/>
				)}
				<ButtonLogin />
			</div>
		</>
	);
};

export default HomeContainer;
