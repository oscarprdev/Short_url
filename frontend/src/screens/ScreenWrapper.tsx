import { ReactNode } from 'react';
import Background from '../components/Background';

const ScreenWrapper = ({ children }: { children: ReactNode }) => {
	return (
		<main className='relative flex flex-col items-center h-screen w-screen overflow-hidden bg-black'>
			{children}
			<Background />
		</main>
	);
};

export default ScreenWrapper;
