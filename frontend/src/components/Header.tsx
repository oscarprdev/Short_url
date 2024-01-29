import { useGlobalStore } from '../store/globalState';
import ButtonLogin from './ButtonLogin';
import ButtonLogout from './ButtonLogout';

const Header = () => {
	const { user } = useGlobalStore();

	return (
		<header className='z-20 flex items-center justify-between py-5 px-10 w-full'>
			{user ? (
				<>
					<h1 className='text-[1.8rem] lg:text-[2rem] font-extrabold'>Short - it</h1>
					<ButtonLogout />
				</>
			) : (
				<ButtonLogin />
			)}
		</header>
	);
};

export default Header;
