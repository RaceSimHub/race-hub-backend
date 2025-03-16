const passwordInput = document.getElementById('password');
const passwordBar = document.getElementById('password-bar');
const passwordMessage = document.getElementById('password-message');
const passwordStrengthDiv = document.getElementById('password-strength');
const confirmPasswordInput = document.getElementById('confirm_password');
const confirmPasswordMessage = document.getElementById('confirm-password-message');
const submitButton = document.getElementById('submit-button');
const nameInput = document.getElementById('name');
const emailInput = document.getElementById('email');

function checkPasswordStrength() {
    if (passwordInput.value === '') {
        passwordStrengthDiv.classList.add('hidden');
        passwordMessage.textContent = '';
        return false;
    } else {
        passwordStrengthDiv.classList.remove('hidden');
    }

    const result = zxcvbn(passwordInput.value);
    const strength = result.score;

    const colors = ['bg-red-500', 'bg-orange-500', 'bg-yellow-500', 'bg-green-500', 'bg-green-700'];
    const messages = [
        'Muito fraca',
        'Fraca',
        'Razoável',
        'Forte',
        'Muito forte'
    ];

    passwordBar.className = `h-full w-${(strength + 1) * 16} ${colors[strength]} rounded transition-all duration-300`;
    passwordMessage.textContent = messages[strength];

    return strength > 0;
}

function checkPasswordMatch() {
    if (passwordInput.value !== confirmPasswordInput.value) {
        confirmPasswordMessage.classList.remove('hidden');
        return false;
    } else {
        confirmPasswordMessage.classList.add('hidden');
        return true;
    }
}

function checkFormValidity() {
    const isPasswordValid = checkPasswordStrength();
    const arePasswordsMatching = checkPasswordMatch();
    const areFieldsFilled = nameInput.value.trim() !== '' && emailInput.value.trim() !== '' && passwordInput.value.trim() !== '' && confirmPasswordInput.value.trim() !== '';

    if (isPasswordValid && arePasswordsMatching && areFieldsFilled) {
        submitButton.disabled = false;
        submitButton.classList.remove('disabled:opacity-50');
    } else {
        submitButton.disabled = true;
        submitButton.classList.add('disabled:opacity-50');
    }
}

passwordInput.addEventListener('input', checkFormValidity);
confirmPasswordInput.addEventListener('input', checkFormValidity);
nameInput.addEventListener('input', checkFormValidity);
emailInput.addEventListener('input', checkFormValidity);