import { IconHeartFilled } from '@tabler/icons-react';

const Footer = () => {
	return (
		<footer className='p-2 z-10 opacity-70'>
			<p className='text-stone-400 w-full flex gap-2'>
				Made with <IconHeartFilled className='text-[var(--contrast)] mx-1' /> by
				<a
					href='https://github.com/oscarprdev'
					target='_blank'
					className='hover:text-[var(--contrast)] duration-300'>
					Oscarprdev
				</a>
			</p>
		</footer>
	);
};

export default Footer;
