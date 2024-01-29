import { ReactNode } from 'react';
import { useUserLogged } from '../hooks/useUserLogged';
import { Redirect } from 'wouter';

const AuthScreen = ({ children, userId }: { children: ReactNode; userId: string }) => {
	const { isLoading, isError } = useUserLogged(userId);

	if (isLoading) {
		return <p>Loading..</p>;
	}

	if (isError) {
		console.log(isError);
		return <Redirect to='/' />;
	}

	return <>{children}</>;
};

export default AuthScreen;
