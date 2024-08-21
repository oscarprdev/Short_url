import { IconLoader2 } from '@tabler/icons-react';

const RedirectingModal = () => {
	return (
		<article className="w-[300px] flex flex-col items-center gap-5 pb-8">
			<p>Redirecting...</p>
			<IconLoader2 data-testid="icon-loader" className="animate-spin" />
		</article>
	);
};

export default RedirectingModal;
