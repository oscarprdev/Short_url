import { IconUser } from '@tabler/icons-react';
import AuthInput from './AuthInput';

const ButtonLogout = () => {
	return (
		<AuthInput
			href='http://localhost:8080/auth/logout'
			contrast>
			<IconUser />
			Logout
		</AuthInput>
	);
};

export default ButtonLogout;
