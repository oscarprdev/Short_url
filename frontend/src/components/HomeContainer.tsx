import { useShortUrl } from '../hooks/useShortUrl';
import { useUrlStore } from '../store/urlStore/urlStore';
import HomeForm from './HomeForm';

const HomeContainer = () => {
	const { url } = useUrlStore();
	const { mutate: addurl } = useShortUrl();

	return (
		<>
			<h1 className='text-[6rem] font-extrabold'>Short - it</h1>
			<div className='flex flex-col items-center gap-5'>
				<h2 className='text-[1.2rem]'>The simplest URL Shortner you were waiting for</h2>
				<HomeForm addUrl={addurl} />
			</div>
			<p>{url}</p>
		</>
	);
};

export default HomeContainer;
