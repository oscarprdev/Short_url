import { useUserLogged } from '../hooks/useUserLogged';
import { IconLoader2 } from '@tabler/icons-react';
import { ReactNode } from 'react';
import { Redirect } from 'wouter';

const AuthScreen = ({ children, userId }: { children: ReactNode; userId: string }) => {
	const { isLoading, isError } = useUserLogged(userId);

	if (isLoading) {
		return (
			<section className="grid place-items-center w-screen h-screen">
				<IconLoader2 className="animate-spin text-contrast" size={30} />
			</section>
		);
	}

	if (isError) {
		return <Redirect to="/" />;
	}

	return <>{children}</>;
};

export default AuthScreen;
