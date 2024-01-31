import { IconHeartFilled } from '@tabler/icons-react';

const Footer = () => {
	return (
		<footer className='absolute bottom-0 w-screen flex justify-center items-center p-2 z-10 opacity-70 h-10'>
			<p className='text-stone-400 flex gap-2'>
				Made with{' '}
				<IconHeartFilled
					data-testid='heart-icon'
					className='text-[var(--contrast)] mx-1'
				/>{' '}
				by
				<a
					href='https://github.com/oscarprdev'
					target='_blank'
					className='hover:text-[var(--contrast)] duration-300'>
					@oscarprdev
				</a>
			</p>
		</footer>
	);
};

export default Footer;
