import { useParams } from 'wouter';
import { useRedirectTo } from '../hooks/useRedirectTo';

const RedirectComponent = () => {
	const { id } = useParams();
	const path = useRedirectTo(id);

	console.log(path);

	window.location.href = path;

	return (
		<section className='flex flex-col h-full items-center justify-center z-10'>
			<p>Redirecting ...</p>
		</section>
	);
};

export default RedirectComponent;
