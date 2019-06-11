import telebot
bot = telebot.TeleBot("845099949:AAHGgsnVDhr0KUqIUW4ilHNXvcN9vXv3hfE")
@bot.message_handler(commands=['start','help'])
def send_flex(message):
	bot.reply_to(message, "Flex")
