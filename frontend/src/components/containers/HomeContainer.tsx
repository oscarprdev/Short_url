import { useShortUrl } from '../../hooks/useShortUrl';
import { useGlobalStore } from '../../store/globalState';
import ButtonLogin from '../ButtonLogin';
import ShortUrlForm from '../ShortUrlForm';
import UrlCardList from '../UrlCardList';

const HomeContainer = () => {
	const { urls } = useGlobalStore();
	const { mutate: addUrl } = useShortUrl();

	return (
		<main className='flex flex-col h-full items-center gap-20 z-10 mt-20'>
			<h1 className='text-[3rem] lg:text-[6rem] font-geistUltra'>Short - it</h1>
			<div className='flex flex-col items-center gap-5'>
				<h2 className='text-[1.6rem] text-center lg:text-[1.2rem]'>The simplest URL Shortner you were waiting for.</h2>
				<ShortUrlForm addUrl={addUrl} />
				{urls && (
					<UrlCardList
						urls={urls}
						isHome
					/>
				)}
				<ButtonLogin />
			</div>
		</main>
	);
};

export default HomeContainer;
