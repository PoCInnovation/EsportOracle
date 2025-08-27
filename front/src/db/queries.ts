import { supabase } from "./supabase";

export const AddOrCreateNewUser = async (walletAddress: string) => {
    try  {

    const { data: existingUser} = await supabase
      .from('users')
      .select('*')
      .eq('wallet_address', walletAddress)
      .single()
    if (existingUser) {
      console.log('User finding:', existingUser)
      return existingUser
    }

    const { data: newUser, error: insertError } = await supabase
      .from('users')
      .insert([
        { wallet_address: walletAddress }
      ])
      .select()
      .single()

    if (insertError) throw insertError

    await supabase
      .from('user_preferences')
      .insert([
        { 
          user_id: newUser.id,
          notifications_enabled: true,
          favorite_teams: []
        }
      ])
    console.log('New user:', newUser)
    return newUser
    } catch (err) {
        console.error('Error creating/retrieving user:', err)
        throw err
    }
}

export const updateUserPseudo = async (walletAddress: string, nameUser: string) => {
    const name = await supabase
    .from('users')
    .insert([
        {pseudo: nameUser}
    ])
    .eq('wallet_address', walletAddress)
    .select()
    .single()
    return name
}