import { useShortUrl } from '../hooks/useShortUrl';
import { useUrlStore } from '../store/urlStore/urlStore';
import HomeForm from './HomeForm';
import HomeUrlCard from './HomeUrlCard';
import { Button } from './ui/button';

const HomeContainer = () => {
	const { url } = useUrlStore();
	const { mutate: addurl } = useShortUrl();

	return (
		<>
			<h1 className='text-[6rem] font-geistUltra'>Short - it</h1>
			<div className='flex flex-col items-center gap-5'>
				<h2 className='text-[1.2rem]'>The simplest URL Shortner you were waiting for.</h2>
				<HomeForm addUrl={addurl} />
				{url && <HomeUrlCard url={url} />}
			</div>
			<Button className='absolute top-8 right-10'>Login</Button>
		</>
	);
};

export default HomeContainer;
