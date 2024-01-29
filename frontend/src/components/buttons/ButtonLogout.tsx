import { IconUser } from '@tabler/icons-react';
import { Button } from '../ui/button';
import { useGlobalStore } from '../../store/globalState';

const ButtonLogout = () => {
	const { clearStore } = useGlobalStore();

	const onLogoutClick = () => {
		clearStore();
		window.location.href = 'http://localhost:8080/auth/logout';
	};
	return (
		<Button onClick={onLogoutClick}>
			<IconUser />
			Logout
		</Button>
	);
};

export default ButtonLogout;
